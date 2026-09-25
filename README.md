# DevOps Project

Клиент-серверное REST API приложение для получения прогноза погоды, развёрнутое в Kubernetes с полным CI/CD-процессом на основе GitOps.

## Описание

Приложение на Go принимает HTTP-запросы и обращается к внешнему Weather API (Visual Crossing) для получения данных о погоде. Демонстрирует полный жизненный цикл: от разработки до автоматической доставки в Kubernetes.

## Стек технологий

- Go
- Docker (multi-stage build), Docker Compose
- Kubernetes (kubeadm), CRI-O, Calico
- Ansible
- Helm
- GitLab CI
- Jenkins (CI)
- Argo CD, Argo CD Image Updater (CD)
- NGINX Ingress Controller
- MetalLB
- Sealed Secrets

## API

### GET /info

Информация о сервисе.

```bash
curl https://weather-app.local/info
json
{
  "author": "m.sveshnikov1",
  "service": "weather",
  "version": "1.0.0"
}
```

## GET /info/weather

Статистика температуры для города и периода.

Параметры:

- `city` (обязательный)
- `date_from` (опционально, YYYY-MM-DD)
- `date_to` (опционально, YYYY-MM-DD)

```bash
curl "https://weather-app.local/info/weather?city=Saint-Petersburg&date_from=2024-03-25&date_to=2024-03-26"
json
{
  "service": "weather",
  "data": {
    "temperature_c": {
      "average": 2.25,
      "median": 2.25,
      "min": -1.2,
      "max": 8
    }
  }
}
```

## Переменные окружения

| Переменная | По умолчанию | Описание |
|------------|--------------|----------|
| `PORT` | `8000` | Порт приложения |
| `VERSION` | `1.0.0` | Версия |
| `AUTHOR` | `m.sveshnikov1` | Автор |
| `API_KEY` | — | Ключ Weather API (обязательный) |

## Развёртывание

### Локально

```bash
docker-compose up --build
```

### Kubernetes

Автоматически через Argo CD при пуше в `master`. Вручную:

```bash
helm upgrade --install weather-app ./weather-app --namespace weather-app
```

## CI/CD

1. Разработчик пушит изменения в Git.
2. Jenkins собирает образ, запускает тесты, публикует в Docker Hub.
3. Argo CD Image Updater обновляет тег образа в Application.
4. Argo CD синхронизирует кластер с Git-репозиторием.

## Отказоустойчивость

- 2 реплики приложения с `podAntiAffinity`.
- `RollingUpdate` (`maxUnavailable: 1`, `maxSurge: 0`).
- Liveness и Readiness пробы.
- PodDisruptionBudget (`minAvailable: 1`).
- ResourceQuota.
- 2 реплики Ingress-контроллера с `podAntiAffinity`.

## Безопасность

- Sealed Secrets для API_KEY.
- TLS-терминация на Ingress (самоподписанный сертификат).
- Отдельный namespace `weather-app`.
