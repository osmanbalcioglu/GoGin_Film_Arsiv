document.getElementById('createForm').addEventListener('submit', function(e) {
    e.preventDefault();
    const data = {
        title: document.getElementById('title').value,
        genre: document.getElementById('genre').value,
        director: document.getElementById('director').value,
        year: parseInt(document.getElementById('year').value),
        rating: parseFloat(document.getElementById('rating').value),
        poster: document.getElementById('poster').value
    };

    fetch('/movies', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    }).then(res => { if(res.ok) window.location.href = "/"; });
});