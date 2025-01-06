class Book:
    def __init__(self, title, writer, cost, stock):
        self.title = title
        self.writer = writer
        self.cost = cost
        self.stock = stock

    def display_details(self):
        return f"Title: {self.title}, Writer: {self.writer}, Cost: {self.cost}, Stock: {self.stock}"

class BookManager:
    def __init__(self, book_inventory=None):
        self.book_inventory = book_inventory if book_inventory else []

    def register_book(self, title, writer, cost, stock):
        try:
            cost = float(cost)
            stock = int(stock)
            if cost < 0 or stock < 0:
                raise ValueError("Cost and stock must be positive numbers.")
            book = Book(title, writer, cost, stock)
            self.book_inventory.append(book)
            print("The book has been added to the inventory!")
        except ValueError as e:
            print(f"Error: {e}")

    def show_books(self):
        if not self.book_inventory:
            print("Currently, there are no books in stock.")
        else:
            for book in self.book_inventory:
                print(book.display_details())

    def find_book_by_title_or_writer(self, search_term):
        found_books = [
            book for book in self.book_inventory 
            if search_term.lower() in book.title.lower() or search_term.lower() in book.writer.lower()
        ]
        if not found_books:
            print("No matching books found.")
        else:
            for book in found_books:
                print(book.display_details())
