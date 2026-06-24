## Sinematik - Go Gin & SQLite Film Arşivi Projesi

Bu proje, daha önce **.NET (C#) Web API** mimarisiyle production standartlarında geliştirilmiş olan Film Arşivi uygulamasının, **Go (Golang) ve Gin Framework** ekosistemine taşınarak asenkron ve modüler mikro mimari prensiplerine uygun şekilde yeniden yapılandırılması sürecini kapsamaktadır. Proje; bağımsız bir veri katmanı (SQLite), runtime optimizasyonu yapılmış arama indeksleri ve canlı metrik analizi sunan entegre bir yönetim paneli (Dashboard) içeren, veri bütünlüğü yüksek kurumsal standartlarda bir **CRUD** uygulamasıdır.

---

## Go Gin'in .NET Ekosistemine Kıyasla Teknik Avantajları

Projenin Go mimarisine taşınması sürecinde, altyapı ve kaynak yönetimi seviyesinde elde edilen majör mühendislik kazanımları şunlardır:

* **Yüksek Verimlilikte Bellek Yönetimi (Low Memory Footprint):** .NET Web API bileşenleri çalışma zamanında (idle) ortalama 100-150 MB RAM blokajı oluştururken, Go derlemesi aynı yük altında yalnızca **10-20 MB RAM** eşiğinde çalışarak kaynak verimliliğini maksimize etmiştir.
* **Bağımsız Dağıtım Mimarisi (Single Binary Deployment):** `go build` çıktısı olarak projenin tüm runtime bağımlılıkları tek bir statik binary (`.exe`) dosyasına indirgenmiştir. Bu sayede hedef sunucu ortamlarında harici bir .NET Runtime veya CLR katmanına ihtiyaç duyulmadan sıfır konfigürasyon (Zero-Dependency Deployment) sağlanmıştır.
* **Optimize Edilmiş İstek Hat Durumu (Execution Performance):** Go, kod bloğunu doğrudan makine diline derleyerek ara katman (IL/JIT compilation) süreçlerini deaktive eder. Bu mimari tercih, HTTP istek-yanıt döngülerindeki gecikme (latency) sürelerini milisaniye altı seviyelere indirgemiştir.
* **Deklaratif ve Şeffaf Katman Yönetimi (Explicit Over Implicit):** .NET ekosisteminin getirdiği soyutlanmış, soyut bağımlılık enjeksiyonları (Implicit DI) veya runtime seviyesinde işleyen karmaşık middleware boru hatları yerine; Go'nun yalın, izlenebilir ve genişletilebilir deterministik yapısı entegre edilmiştir.

---

## Geliştirme Sürecinde Karşılaşılan Zorluklar, Teknik Bariyerler ve Çözümler

Yüksek soyutlama düzeyine sahip kurumsal .NET mimarisinden, Go'nun minimalist ve performans odaklı yapısına geçişte karşılaşılan teknik bariyerler ve uygulanan mühendislik çözümleri:

### 1. Monolitik Yapının Kırılması ve Katmanlı Mimari (SoC) Tasarımı
* **Zorluk:** İlk prototip aşamasında iş mantığının ve yönlendirmelerin tek bir `main.go` üzerinde konumlandırılması, kodun sürdürülebilirliğini (maintainability) ve test edilebilirliğini düşürmekteydi.
* **Çözüm:** **MVC (Model-View-Controller)** mimari deseni projenin doğasına uygun şekilde adapte edilmiştir; veri modelleri `models`, iş mantığı ve API uç noktaları `controllers`, arayüz şablonları ise `views` katmanına izole edilerek katmanlı mimari (Layered Architecture) standartları sağlanmıştır.

### 2. İşletim Sistemi Seviyesinde Şablon Yükleme Çakışmaları (`undefined` Hatası)
* **Hata:** `html/template: "index.html" is undefined`
* **Analiz:** Gin Framework üzerindeki `LoadHTMLGlob("views/**/*")` fonksiyonunun, Windows işletim sistemi dosya yolu ayırıcıları (`\`) nedeniyle alt klasör hiyerarşisinde (`layouts`) adlandırma çakışması yaşadığı ve şablon bellek haritasını (template cache map) bozduğu tespit edilmiştir.
* **Çözüm:** Dinamik glob deseni deaktive edilerek, `r.LoadHTMLFiles()` metodu vasıtasıyla tüm HTML bileşenleri ve master-page şablonları deterministik olarak çalışma zamanına enjekte edilmiş ve şablon yükleme riskleri ortadan kaldırılmıştır.

### 3. Veri Katmanında Case-Sensitivity (Büyük/Küçük Harf) Yönetimi
* **Zorluk:** SQLite veritabanı motoru varsayılan olarak `LIKE` sorgularında ve UTF-8 karakter setlerinde büyük/küçük harf duyarlı (Case-Sensitive) çalışmaktadır. Bu durum, arama endeksinin sorgu isabet oranını azaltmaktaydı.
* **Çözüm:** GORM veri katmanı üzerinde ham SQL manipülasyonu (Raw SQL Query) yapılarak hem veritabanı kolonları hem de istemciden gelen arama parametreleri `LOWER()` fonksiyonu ile normalize edilmiştir (`LOWER(title) LIKE LOWER(?)`). Böylece harf duyarsız (Case-Insensitive) arama mimarisi kurulmuştur.

### 4. TCP Soket Blokajları ve Süreç Yönetimi (`Port Already in Use`)
* **Hata:** Geliştirme ve hot-reload senaryolarında, ilgili ağ portunun önceki oturumlar tarafından serbest bırakılmaması.
* **Analiz:** Windows çekirdeğinin TCP soketini `TIME_WAIT` durumunda tutmasından kaynaklanan soket blokajları tespit edilmiştir.
* **Çözüm:** Lokal script süreçlerine ve CI/CD pipeline yapılarına `taskkill /f /im main.exe` komut zinciri entegre edilerek soketlerin zorunlu olarak temizlenmesi (socket reuse) ve kararlı başlatma süreçleri otomatize edilmiştir.

---

## Proje Klasör Yapısı

Proje, Sorumlulukların Ayrılması (**Separation of Concerns - SoC**) prensibine uygun olarak aşağıdaki kurumsal dizin mimarisinde kurgulanmıştır:

```text
FilmArsivGo/
│
├── controllers/          # Uygulama İş Mantığı, API Yönetimi ve Sayfa Render İşlemleri
│   └── film_controller.go
│
├── models/               # Veri Modelleri (Şemalar) ve GORM Veritabanı Bağlantısı
│   └── film.go
│
├── public/               # Statik Varlıklar (Static Assets)
│   ├── css/
│   │   └── style.css     # UI Bileşen Stilleri
│   └── js/
│       ├── api.js        # Global AJAX/Fetch Silme (DELETE) Operasyonları
│       ├── create.js     # Kayıt (POST) İstemci Mantığı
│       └── update.js     # Güncelleme (PUT) Mantığı & Dinamik Poster Önizleme İşleyicisi
│
├── views/                # HTML5 / Go Template Arayüz Bileşenleri
│   ├── layouts/
│   │   ├── header.html   # Master Page Üst Bölüm ve CSS Bağımlılıkları
│   │   └── footer.html   # Master Page Alt Bölüm ve Kapanış Etiketleri
│   ├── index.html        # Veritabanı Tabanlı Gelişmiş Filtreleme & Dashboard Ekranı
│   ├── create.html       # Yeni Veri Giriş Formu
│   └── update.html       # Veri Güncelleme Formu (Real-time Önizleme Destekli)
│
├── main.go               # Uygulama Giriş Noktası, Rota Tanımlamaları ve Soket Dinleyici
├── go.mod                # Go Bağımlılık ve Modül Bildirimi
└── go.sum                # Bağımlılık Doğrulama ve Güvenlik İmzaları

## Projeyi Yerelde Çalıştırma Talimatı
Uygulamanın yerel geliştirme ortamında (localhost) ayağa kaldırılması için aşağıdaki adımların terminal üzerinden yürütülmesi gerekmektedir:

1. Port Temizliği ve Mevcut Proseslerin Sonlandırılması
Daha önceki oturumlardan kalan hayalet proseslerin (zombie processes) ağ soketini işgal etmesini engellemek adına ilgili portu zorunlu olarak boşa çıkarın:

taskkill /f /im main.exe
2. Uygulama Bağımlılıklarının Doğrulanması ve Çalıştırılması
Projenin kök dizininde (root) bağımlılıkların indirilmesini sağlayarak uygulamayı derleyin ve başlatın:

go run main.go
3. İstemci Erişimi
Sunucu başarıyla başlatıldıktan sonra HTTP çoklu sayfa yönlendirmelerini test etmek için tarayıcı üzerinden aşağıdaki uç noktaya istek atın:

http://localhost:8085/
##Kullanılan Teknolojiler
Arka Plan (Backend) & Veri Katmanı
Dil: Go (Golang) v1.22+ (Yüksek performanslı, concurrency odaklı derleme dili)

Web Framework: Gin Gonic v1.10+ (Minimalist, yüksek performanslı HTTP Router kütüphanesi)

ORM: GORM (Go Object Relational Mapping - SQLite Sürücüsü ile)

Veritabanı Motoru: SQLite3 (Gömülü, dosya tabanlı ve ACID uyumlu ilişkisel veritabanı)

Ön Yüz (Frontend) & Asenkron Yönetim
Arayüz Şablon Motoru: Go html/template (Güvenli, XSS enjeksiyon korumalı backend-render mimarisi)

CSS Framework: Bootstrap v5.3 (Responsive grid ve UI bileşen kütüphanesi)

İstemci Mantığı: Vanilla JavaScript (ES6+ standartlarında asenkron event yönetimi)

Veri İletişim Protokolü: Fetch API (JSON tabanlı RESTful asenkron istek yönetimi)

⭐ Bu geçiş çalışması; .NET ekosisteminin getirdiği yüksek soyutlama katmanlarının ötesine geçerek, bellek optimizasyonu, ağ soket yönetimi ve derleme mimarileri üzerinde alt seviye (low-level) sistem vizyonu ve mühendislik yetkinliği kazanmak amacıyla kurgulanmıştır.
