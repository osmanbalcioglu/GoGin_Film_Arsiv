const urlParams = new URLSearchParams(window.location.search);
const movieId = urlParams.get('id');

if (!movieId) {
    alert("Geçersiz film ID!");
    window.location.href = "/";
}

fetch('/movies/' + movieId)
.then(res => res.json())
.then(film => {
    document.getElementById('title').value = film.title;
    document.getElementById('genre').value = film.genre;
    document.getElementById('director').value = film.director;
    document.getElementById('year').value = film.year;
    document.getElementById('rating').value = film.rating;
    
    const posterInput = document.getElementById('poster');
    posterInput.value = film.poster;

    const previewContainer = document.getElementById('posterPreviewContainer');
    const previewImg = document.getElementById('posterPreview');
    if (film.poster) {
        previewImg.src = film.poster;
        previewContainer.classList.remove('d-none');
    }
});

document.getElementById('poster').addEventListener('input', function(e) {
    const previewContainer = document.getElementById('posterPreviewContainer');
    const previewImg = document.getElementById('posterPreview');
    
    if (e.target.value.trim() !== "") {
        previewImg.src = e.target.value;
        previewContainer.classList.remove('d-none');
    } else {
        previewContainer.classList.add('d-none');
    }
});

document.getElementById('updateForm').addEventListener('submit', function(e) {
    e.preventDefault();
    const data = {
        title: document.getElementById('title').value,
        genre: document.getElementById('genre').value,
        director: document.getElementById('director').value,
        year: parseInt(document.getElementById('year').value),
        rating: parseFloat(document.getElementById('rating').value),
        poster: document.getElementById('poster').value
    };

    fetch('/movies/' + movieId, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data)
    }).then(res => { if(res.ok) window.location.href = "/"; });
});