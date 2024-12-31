from flask import Flask, request, jsonify

app = Flask(__name__)

class BorrowLimitExceededException(Exception):
    pass

class Book:
    def __init__(self, book_id, title, author, type_):
        self.book_id = book_id
        self.title = title
        self.author = author
        self.type = type_
        self.is_borrowed = False

    def borrow(self):
        if self.is_borrowed:
            raise Exception("Book is already borrowed")
        self.is_borrowed = True

    def return_book(self):
        self.is_borrowed = False

class Member:
    def __init__(self, member_id, name, membership_type):
        self.member_id = member_id
        self.name = name
        self.membership_type = membership_type
        self.borrowed_books = []

    def borrow_book(self, book):
        if len(self.borrowed_books) >= self.max_books():
            raise BorrowLimitExceededException("Borrowing limit exceeded!")
        book.borrow()
        self.borrowed_books.append(book)

    def return_book(self, book):
        book.return_book()
        self.borrowed_books.remove(book)

    def max_books(self):
        return 5 if self.membership_type == 'Premium' else 3

class Library:
    def __init__(self):
        self.books = []
        self.members = []

    def add_book(self, title, author, type_):
        book_id = len(self.books) + 1
        book = Book(book_id, title, author, type_)
        self.books.append(book)
        return book

    def add_member(self, name, membership_type):
        member_id = len(self.members) + 1
        member = Member(member_id, name, membership_type)
        self.members.append(member)
        return member

library = Library()

@app.route('/add_book', methods=['POST'])
def add_book():
    data = request.get_json()
    book = library.add_book(data['title'], data['author'], data['type'])
    return jsonify({"message": f"Book '{book.title}' added!"})

@app.route('/add_member', methods=['POST'])
def add_member():
    data = request.get_json()
    member = library.add_member(data['name'], data['type'])
    return jsonify({"message": f"Member '{member.name}' registered!"})

if __name__ == "__main__":
    app.run(debug=True)
