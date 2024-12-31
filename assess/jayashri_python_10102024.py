""" Case study:

Books:
Attributes: Unique book ID, title, author, and availability status (available/borrowed).
Methods:
Add new books to the collection.
Update book status when borrowed or returned.
Types: Categorized as fiction or non-fiction.

Members:
Attributes: Unique member ID, name, list of borrowed books, and borrowing limit based on membership type (Regular: 3 books, Premium: 5 books).
Methods:
Borrow a book if available.
Return a borrowed book and update records.
Types: Regular and Premium members with different borrowing privileges.

Library:
Attributes: Collection of books and list of members.
Methods:
Add books to the collection.
Register new members.
Handle book borrowing and returns, including updates to book status and member records."""







class BorrowLimitExceededException(Exception):
    def __init__(self, message):
        self.message = message
        super().__init__(self.message)


class Book:
    def __init__(self, book_id, title, author, status="available"):
        self.book_id = book_id
        self.title = title
        self.author = author
        self.status = status

    def borrow(self):
        if self.status == "available":
            self.status = "borrowed"
        else:
            raise ValueError(f"Book '{self.title}' is already borrowed.")

    def return_book(self):
        if self.status == "borrowed":
            self.status = "available"
        else:
            raise ValueError(f"Book '{self.title}' is not borrowed.")


class Member:
    def __init__(self, member_id, name, max_books):
        self.member_id = member_id
        self.name = name
        self.borrowed_books = []
        self.max_books = max_books

    def borrow_book(self, book):
        if len(self.borrowed_books) < self.max_books:
            if book.status == "available":
                book.borrow()
                self.borrowed_books.append(book)
            else:
                raise ValueError(f"Book '{book.title}' is currently unavailable.")
        else:
            raise BorrowLimitExceededException(f"{self.name} has reached the borrowing limit of {self.max_books} books.")

    def return_book(self, book):
        if book in self.borrowed_books:
            book.return_book()
            self.borrowed_books.remove(book)
        else:
            raise ValueError(f"{self.name} hasn't borrowed the book '{book.title}'.")

class RegularMember(Member):
    def __init__(self, member_id, name):
        super().__init__(member_id, name, max_books=3)


class PremiumMember(Member):
    def __init__(self, member_id, name):
        super().__init__(member_id, name, max_books=5)


class Library:
    def __init__(self):
        self.book_collection = []
        self.members = []

    def add_book(self, book):
        if not book.title or not book.author:
            raise ValueError("Book must have a title and an author.")
        self.book_collection.append(book)

    def register_member(self, member):
        if not member.name:
            raise ValueError("Member must have a name.")
        self.members.append(member)

    def lend_book(self, member, book):
        try:
            member.borrow_book(book)
            print(f"Book '{book.title}' borrowed by {member.name}.")
        except (ValueError, BorrowLimitExceededException) as e:
            print(f"Error: {e}")

    def receive_return(self, member, book):
        try:
            member.return_book(book)
            print(f"Book '{book.title}' returned by {member.name}.")
        except ValueError as e:
            print(f"Error: {e}")


if __name__ == "__main__":
    library = Library()


    book1 = Book(book_id=1, title="1984", author="George Orwell")
    book2 = Book(book_id=2, title="Sapiens", author="Yuval Noah Harari")
    library.add_book(book1)
    library.add_book(book2)


    regular_member = RegularMember(member_id=101, name="Alice")
    premium_member = PremiumMember(member_id=102, name="Bob")
    library.register_member(regular_member)
    library.register_member(premium_member)


    library.lend_book(regular_member, book1)
    library.lend_book(premium_member, book2)

  
    library.receive_return(regular_member, book1)
    library.receive_return(premium_member, book2)

   
    try:
        library.lend_book(regular_member, book2)
        library.lend_book(regular_member, book1) 
    except BorrowLimitExceededException as e:
        print(f"Exception: {e.message}")



book3 = Book(book_id=3, title="The Catcher in the Rye", author="J.D. Salinger")
book4 = Book(book_id=4, title="The Great Gatsby", author="F. Scott Fitzgerald")
library.add_book(book3)
library.add_book(book4)


library.lend_book(regular_member, book3)  
library.lend_book(regular_member, book4)  
