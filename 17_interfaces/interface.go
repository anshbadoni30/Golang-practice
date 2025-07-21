package main
import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct{
	gateway paymenter
}

func (p payment) makepayment(amount float32){
	p.gateway.pay(amount)
}

type razorpay struct{}
func (r razorpay)pay(amount float32){
	fmt.Println("amount to be paid by razorpay is ",amount)
}

type paypal struct{}
func (p paypal)pay(amount float32){
	fmt.Println("amount to be paid by paypal is ",amount)
}

func main(){
	//razor1:=razorpay{}
	paypal1:=paypal {}
	payment1:=payment{
		gateway: paypal1,
	}
	payment1.makepayment(1000.0)

}