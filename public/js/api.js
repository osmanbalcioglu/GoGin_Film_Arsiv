function deleteFilm(id) {
    if(confirm('Bu filmi arşivden silmek istediğinize emin misiniz?')) {
        fetch('/movies/' + id, { method: 'DELETE' })
        .then(res => { if(res.ok) window.location.reload(); });
    }
}