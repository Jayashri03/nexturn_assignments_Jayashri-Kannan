document.getElementById("addBookForm").addEventListener("submit", function (event) {
    event.preventDefault();
    const title = document.getElementById("bookTitle").value;
    const author = document.getElementById("bookAuthor").value;
    const type = document.getElementById("bookType").value;

    fetch('http://localhost:5000/add_book', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({title, author, type})
    }).then(response => response.json())
      .then(data => alert(data.message));
});

document.getElementById("addMemberForm").addEventListener("submit", function (event) {
    event.preventDefault();
    const name = document.getElementById("memberName").value;
    const type = document.getElementById("membershipType").value;

    fetch('http://localhost:5000/add_member', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({name, type})
    }).then(response => response.json())
      .then(data => alert(data.message));
});
