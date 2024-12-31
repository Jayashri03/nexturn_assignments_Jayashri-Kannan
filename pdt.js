class product{
    #price;
    #quantity;
    name;
    constructor(Name, Price, Quantity) {
        this.name = Name;
        this.#price = Price;
        this.#quantity = Quantity || 0;
    }
    calculateTotal(){
        console.log("The total pricce of product is: "+ this.#price*this.#quantity);
    }
    updateQuantity(newQuantity){
        if(newQuantity<0){
            console.log("Enter a positive number.");
            return;
        }
        this.#quantity=newQuantity;
        console.log("The quantity is updated to "+this.#quantity);
    }
    displayInfo(){
        console.log("There are "+this.#quantity+" "+ this.name+" of price "+this.#price+" rupees");
    }
}


class DigitalProduct extends product{
    #LisenceKey;
    constructor(Name,price,LisenceKey){
        super(Name,price);
        this.#LisenceKey=LisenceKey;
    }
    displayInfo(){
        console.log(this.#LisenceKey,this.name)
    }
    download(){
        console.log("Downloading the product with lisence key: "+ this.#LisenceKey)
    }
}
var prod = new DigitalProduct("apple",300,"2ABSCDE");
prod.displayInfo();
prod.download();