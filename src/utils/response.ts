export type Response<T extends any = {}> = {
    code: number,
    type?: string,
    message?: string,
    error?: string,
    extra?: T
}