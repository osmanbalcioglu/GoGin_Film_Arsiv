function calculateStats() {
    const movies = document.querySelectorAll('.movie-item');
    const totalMovies = movies.length;
    let maxRating = -1; 
    let topMovie = "Yok";

    movies.forEach(movie => {
        const r = parseFloat(movie.getAttribute('data-rating'));
        const t = movie.getAttribute('data-title');
        
        if (!isNaN(r) && r > maxRating) { 
            maxRating = r; 
            topMovie = t; 
        }
    });

    const displayRating = maxRating >= 0 ? maxRating : '-';

    const statsHtml = `
        <div class="col-md-4 mb-3">
            <div class="card p-3 shadow-sm text-center bg-white border-start border-primary border-5">
                <h6 class="text-muted text-uppercase small fw-bold">Toplam Film</h6>
                <h2 class="fw-bold text-dark m-0">${totalMovies}</h2>
            </div>
        </div>
        <div class="col-md-4 mb-3">
            <div class="card p-3 shadow-sm text-center stats-card">
                <h6 class="text-white-50 text-uppercase small fw-bold">En Yüksek Puan</h6>
                <h2 class="fw-bold m-0">${displayRating}</h2>
            </div>
        </div>
        <div class="col-md-4 mb-3">
            <div class="card p-3 shadow-sm text-center bg-white border-start border-success border-5">
                <h6 class="text-muted text-uppercase small fw-bold">Zirvedeki Film</h6>
                <h2 class="fw-bold text-success m-0 fs-5 text-truncate" title="${topMovie}">${topMovie}</h2>
            </div>
        </div>
    `;
    
    const statsPanel = document.getElementById('statsPanel');
    if (statsPanel) {
        statsPanel.innerHTML = statsHtml;
    }
}

window.onload = calculateStats;