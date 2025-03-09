# Todo Planning Application

Bu uygulama, farklı kaynaklardan gelen görevleri toplayıp, developer'lara haftalık olarak atayan bir planlama sistemidir.

## Özellikler

- Farklı kaynaklardan görev toplama (Provider entegrasyonu)
- Developer yönetimi (Seniority bazlı iş dağılımı)
- Haftalık sprint planlama
- Otomatik iş dağıtımı
- PostgreSQL veritabanı entegrasyonu

## Teknolojiler

- Go 1.23
- PostgreSQL 15
- Docker & Docker Compose
- GORM (ORM)
- Wire (Dependency Injection)
- Ginkgo & Gomega (Testing)

## Başlangıç

### Gereksinimler

- Docker ve Docker Compose
- Go 1.23 veya üzeri
- Make

### Kurulum

1. Repoyu klonlayın:
```bash
git clone https://github.com/yourusername/todo-planning.git
cd todo-planning
```

2. Uygulamayı başlatın:
```bash
# Development modunda
docker compose up --build

# Veya local'de çalıştırmak için
make run
```

## Development

### Ortam Hazırlığı

```bash
# Wire kodunu generate et
make wire

# Mock'ları generate et
make generate-mocks
```

### Test

```bash
# Testleri çalıştır
make test

# Veya docker ile tüm testleri çalıştır
make test-composer-up
```

### Proje Yapısı
