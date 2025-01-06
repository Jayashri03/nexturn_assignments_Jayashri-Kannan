from flask import Flask, request, jsonify
import sqlite3

app = Flask(__name__)

def get_database_connection():
    connection = sqlite3.connect('library.db')
    connection.row_factory = sqlite3.Row
    return connection

@app.route('/library/books', methods=['POST'])
def register_book():
    data = request.json
    if not all(key in data for key in ['title', 'writer', 'year_published', 'category']):
        return jsonify({"error": "Incomplete data", "message": "All fields are required"}), 400
    try:
        connection = get_database_connection()
        cursor = connection.cursor()
        cursor.execute(
            'INSERT INTO books (title, writer, year_published, category) VALUES (?, ?, ?, ?)',
            (data['title'], data['writer'], data['year_published'], data['category'])
        )
        connection.commit()
        connection.close()
        return jsonify({"message": "Book registered successfully"}), 201
    except Exception as e:
        return jsonify({"error": "Database error", "message": str(e)}), 500

@app.route('/library/books', methods=['GET'])
def list_books():
    connection = get_database_connection()
    books = connection.execute('SELECT * FROM books').fetchall()
    connection.close()
    return jsonify([dict(book) for book in books])

@app.route('/library/books/<int:id>', methods=['GET'])
def get_book(id):
    connection = get_database_connection()
    book = connection.execute('SELECT * FROM books WHERE id = ?', (id,)).fetchone()
    connection.close()
    if book is None:
        return jsonify({"error": "Book not found"}), 404
    return jsonify(dict(book))

@app.route('/library/books/<int:id>', methods=['PUT'])
def modify_book(id):
    data = request.json
    connection = get_database_connection()
    book = connection.execute('SELECT * FROM books WHERE id = ?', (id,)).fetchone()
    if book is None:
        connection.close()
        return jsonify({"error": "Book not found"}), 404
    connection.execute(
        'UPDATE books SET title = ?, writer = ?, year_published = ?, category = ? WHERE id = ?',
        (data['title'], data['writer'], data['year_published'], data['category'], id)
    )
    connection.commit()
    connection.close()
    return jsonify({"message": "Book details updated successfully"})

@app.route('/library/books/<int:id>', methods=['DELETE'])
def remove_book(id):
    connection = get_database_connection()
    book = connection.execute('SELECT * FROM books WHERE id = ?', (id,)).fetchone()
    if book is None:
        connection.close()
        return jsonify({"error": "Book not found"}), 404
    connection.execute('DELETE FROM books WHERE id = ?', (id,))
    connection.commit()
    connection.close()
    return jsonify({"message": "Book removed successfully"})

if __name__ == '__main__':
    app.run(debug=True)
