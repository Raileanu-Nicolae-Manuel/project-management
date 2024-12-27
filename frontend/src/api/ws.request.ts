export class WSRequest {
    private readonly ws: WebSocket
    constructor(private readonly url: string) {
        this.ws = new WebSocket(url)
    }

    async send(message: string) {
        this.ws.send(message)
    }

    async onMessage(callback: (message: string) => void) {
        this.ws.onmessage = (event) => {
            callback(event.data)
        }
    }

    async close() {
        this.ws.close()
    }

    async onClose(callback: () => void) {
        this.ws.onclose = () => {
            callback()
        }
    }

    async onError(callback: (error: Event) => void) {
        this.ws.onerror = (event) => {
            callback(event)
        }
    }

    async onOpen(callback: () => void) {
        this.ws.onopen = () => {
            callback()
        }
    }

    async onReconnect(callback: () => void) {
        this.ws.addEventListener('reconnect', () => {
            callback()
        })
    }
}