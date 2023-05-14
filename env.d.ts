/// <reference types="vite/client" />

declare module 'element-plus' {
    type MessageType = "info" | "warning" | "success" | "error"

    const ElNotification: (info?: {
        type: MessageType,
        message?: string,
        title?: string
    }) => void

    const ElMessageBox: {
        alert: (message?: string, title?: string, options?: {
            confirmButtonText?: string,
            type?: MessageType
        }) => void
    }

    export { ElNotification, ElMessageBox }
}