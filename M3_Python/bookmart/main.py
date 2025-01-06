from book_management import BookManager
from customer_management import CustomerManager
from sales_management import SalesManager

def main():
    book_manager = BookManager()
    customer_manager = CustomerManager()
    sales_manager = SalesManager()

    while True:
        print("\nWelcome to BookMart!")
        print("1. Book Management")
        print("2. Customer Management")
        print("3. Sales Management")
        print("4. Exit")
        choice = input("Enter your choice: ")

        if choice == "1":
            print("\n1. Register Book")
            print("2. Show Books")
            print("3. Find Book")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                title = input("Title: ")
                writer = input("Writer: ")
                cost = input("Cost: ")
                stock = input("Stock: ")
                book_manager.register_book(title, writer, cost, stock)
            elif sub_choice == "2":
                book_manager.show_books()
            elif sub_choice == "3":
                search_term = input("Enter title or writer to search: ")
                book_manager.find_book_by_title_or_writer(search_term)

        elif choice == "2":
            print("\n1. Register Customer")
            print("2. List Customers")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                full_name = input("Name: ")
                contact_email = input("Email: ")
                phone_number = input("Phone: ")
                customer_manager.register_client(full_name, contact_email, phone_number)
            elif sub_choice == "2":
                customer_manager.list_clients()

        elif choice == "3":
            print("\n1. Sell Book")
            print("2. View Sales")
            sub_choice = input("Enter your choice: ")
            if sub_choice == "1":
                phone_number = input("Enter Customer Phone: ")
                customer = customer_manager.find_client_by_phone(phone_number)
                if customer is None:
                    print("Customer not found!")
                    continue
                book_title = input("Enter Book Title: ")
                book = next((book for book in book_manager.book_inventory if book.title.lower() == book_title.lower()), None)
                if book is None:
                    print("Book not found!")
                    continue
                quantity = input("Quantity: ")
                sales_manager.process_sale(customer, book, quantity)
            elif sub_choice == "2":
                sales_manager.list_sales()

        elif choice == "4":
            print("Thank you for using BookMart!")
            break

        else:
            print("Invalid choice. Please try again!")

if __name__ == "__main__":
    main()
