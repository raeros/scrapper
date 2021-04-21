package scrapper                                                                                                                                                                          
import (
  "golang.org/x/net/html"
  "io"
)

func isTitleElement(n *html.Node) bool {
  return n.Type == html.ElementNode && n.Data == "body"
  //return n.Type == html.ElementNode && n.Data == "title"
}

func traverse(n *html.Node) (string, bool) {

	  if isTitleElement(n) {
       println("___")
       println(n.LastChild.Namespace)
       println("___")
       println(n.LastChild.Data)
       println("___")
       println(n.LastChild.Type)
       println("___")
       println(n.LastChild.DataAtom)
       println("___")
       println(n.LastChild.Attr)
       println("___")

      return n.FirstChild.Data, true
	  }       

    for c := n.FirstChild; c != nil; c = c.NextSibling {                                  
      result, ok := traverse(c)
      if ok {
        return result, ok
      }
    }

    return "", false
}                                                                                          
                                                                                              
func GetHtmlTitle(body io.Reader) (string, bool) {                                            
    body_parsed, err := html.Parse(body)

    if err != nil {                                                                         
      panic("Fail to parse HTML.")
    }

    return traverse(body_parsed)

}  
