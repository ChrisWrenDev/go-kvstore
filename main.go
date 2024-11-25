package main

type Storer interface {
	Put(string, []byte) error
	Get(string )([]byte, error)
    Update(string, []byte) error
    Delete(string)([]byte, error)
}



func main(){

}
