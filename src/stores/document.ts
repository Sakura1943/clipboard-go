import { defineStore } from 'pinia'
import base from '@/utils/base'
import { type Response } from '@/utils/response'

export type PageData = {
    id: number,
    path: string,
    text: string,
    user_name: string
}

export const documentStore = defineStore('docs-list', () => {
    const getDocuments = (token: string, callback: (data: Response<PageData[] | undefined>) => void) => {
        const url = `${base.backendUrl}/api/document/list`
        fetch(url, {
            method: 'GET',
            headers: {
                token
            }
        })
            .then(data => data.json())
            .then(data => callback(data))
            .catch(data => callback(data))
    }

    const getOneDocument = (path: string, lang: string, callback: (data: Response<PageData | undefined>) => void) => {
        const url = `${base.backendUrl}/api/document/search`
        let formData = new FormData()
        formData.append('path', path)
        formData.append('lang', lang)
        fetch(url, {
            method: 'POST',
            body: formData
        }).then(data => data.json())
            .then(data => callback(data))
            .catch(data => callback(data))
    }

    const deleteDocument = (token: string, path: string, callback: (data: Response<{} | undefined>) => void) => {
        path = encodeURI(path)
        const url = `${base.backendUrl}/api/document/delete?path=${path}`
        fetch(url, {
            method: 'DELETE',
            headers: {
                token
            }
        }).then(data => data.json())
            .then(data => callback(data))
            .catch(data => callback(data))
    }
    return { getDocuments, getOneDocument, deleteDocument }
})
