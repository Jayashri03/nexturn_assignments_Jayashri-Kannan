// Book class
class Book {
    constructor(id, title, author) {
        this.book_id = id;
        this.title = title;
        this.author = author;
        this.status = 'available';
    }

    borrow() {
        if (this.status === 'available') {
            this.status = 'borrowed';
        } else {
            throw new Error('Book is already borrowed');
        }
    }

    returnBook() {
        this.status = 'available';
    }
}

// Member class
class Member {
    constructor(id, name, type) {
        this.member_id = id;
        this.name = name;
        this.type = type;
        this.borrowed_books = [];
        this.max_books = type === 'premium' ? 5 : 3;
    }

    borrowBook(book) {
        if (this.borrowed_books.length >= this.max_books) {
            throw new BorrowLimitExceededException('Borrow limit exceeded');
        } else {
            book.borrow();
            this.borrowed_books.push(book);
        }
    }

    returnBook(book) {
        const index = this.borrowed_books.indexOf(book);
        if (index > -1) {
            this.borrowed_books.splice(index, 1);
            book.returnBook();
        } else {
            throw new Error('Book not borrowed by this member');
        }
    }
}

// Library class
class Library {
    constructor() {
        this.book_collection = [];
        this.members = [];
    }

    addBook(title, author) {
        const newBook = new Book(this.book_collection.length + 1, title, author);
        this.book_collection.push(newBook);
        this.updateBookList();
    }

    registerMember(name, type) {
        const newMember = new Member(this.members.length + 1, name, type);
        this.members.push(newMember);
        this.updateMemberList();
    }

    updateBookList() {
        const bookList = document.getElementById('book-list');
        bookList.innerHTML = '';
        this.book_collection.forEach(book => {
            const bookDiv = document.createElement('div');
            bookDiv.innerText = `${book.title} by ${book.author} - ${book.status}`;
            bookList.appendChild(bookDiv);
        });
    }

    updateMemberList() {
        const memberList = document.getElementById('member-list');
        memberList.innerHTML = '';
        this.members.forEach(member => {
            const memberDiv = document.createElement('div');
            memberDiv.innerText = `${member.name} (${member.type}) - Books Borrowed: ${member.borrowed_books.length}`;
            memberList.appendChild(memberDiv);
        });
    }
}

// Custom exception
class BorrowLimitExceededException extends Error {
    constructor(message) {
        super(message);
        this.name = 'BorrowLimitExceededException';
    }
}

// Instantiate the library
const library = new Library();

// Add event listeners for adding books and members
document.getElementById('add-book-form').addEventListener('submit', function(e) {
    e.preventDefault();
    const title = document.getElementById('book-title').value;
    const author = document.getElementById('book-author').value;
    if (title && author) {
        library.addBook(title, author);
    }
});

document.getElementById('add-member-form').addEventListener('submit', function(e) {
    e.preventDefault();
    const name = document.getElementById('member-name').value;
    const type = document.getElementById('member-type').value;
    if (name) {
        library.registerMember(name, type);
    }
});
