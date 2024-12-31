class BankAccount {
    // Private Fields (Encapsulation)
    #accountNumber;
    #accountBalance;
    // Public Fields (Abstraction)
    accountHolderName;
    
    
    // Attributes
    constructor(accNo, accHolderName, accBalance) {
    this.#accountNumber = accNo;
    this.accountHolderName = accHolderName;
    this.#accountBalance = accBalance;
    }
    // Methods
    deposit(amount) {
    this.#accountBalance += amount;
    console.log(amount + " deposited successfully. Current balance is " + this.#accountBalance);
    }
    withdraw(amount) {
    if (this.#accountBalance < amount) {
    console.log("Insufficient balance");
    return;
    }
    this.#accountBalance -= amount;
    console.log(amount + " withdrawn successfully. Current balance is " + this.#accountBalance);
    }
    display() {
    console.log("Account Number: " + this.#accountNumber);
    console.log("Account Holder Name: " + this.accountHolderName);
    console.log("Account Balance: " + this.#accountBalance);
    }
    // Getter Method (Read-Only Property)
    _getAccountBalance() {
    return this.#accountBalance;
    }
    // Setter Method (Write-Only Property)
    _setAccountBalance(amount) {
    this.#accountBalance = amount;
    }
    }
    
    class SavingAccount extends BankAccount {
    // Private Fields
    #interestRate;
    constructor(accNo, accHolderName, accBalance, interestRate) {
    super(accNo, accHolderName, accBalance);
    this.#interestRate = interestRate;
    }
    calculateInterest() {
    let interest = this._getAccountBalance() * this.#interestRate / 100;
    console.log("Interest: " + interest);
    return interest;
    }
    
    applyInterest() {
    let interest = this.calculateInterest();
    this._setAccountBalance(this._getAccountBalance() + interest);
    console.log("Interest applied successfully. Current balance is " + this._getAccountBalance());
    }
    }
    
    var account = new SavingAccount(1234, "John Doe", 5000, 5);
    account.deposit(500); // Output: 500 deposited successfully. Current balance is 5500
    account.withdraw(2000); // Output: 2000 withdrawn successfully. Current balance is 3500
    account.calculateInterest(); // Output: Interest: 175
    account.applyInterest(); // Output: Interest applied successfully. Current balance is 3675