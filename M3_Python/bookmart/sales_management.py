from customer_management import Customer

class Transaction(Customer):
    def __init__(self, full_name, contact_email, phone_number, book_title, quantity_sold):
        super().__init__(full_name, contact_email, phone_number)
        self.book_title = book_title
        self.quantity_sold = quantity_sold

class SalesManager:
    def __init__(self, sales_records=None):
        self.sales_records = sales_records if sales_records else []

    def process_sale(self, customer, book, quantity):
        try:
            quantity = int(quantity)
            if book.stock >= quantity:
                book.stock -= quantity
                transaction = Transaction(customer.full_name, customer.contact_email, customer.phone_number, book.title, quantity)
                self.sales_records.append(transaction)
                print(f"Sale completed! Remaining stock: {book.stock}")
            else:
                print("Insufficient stock available.")
        except ValueError:
            print("Invalid quantity entered!")

    def list_sales(self):
        if not self.sales_records:
            print("No sales records found.")
        else:
            for sale in self.sales_records:
                print(f"Customer: {sale.full_name}, Book: {sale.book_title}, Quantity Sold: {sale.quantity_sold}")
