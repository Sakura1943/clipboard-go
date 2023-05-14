import { createRouter, createWebHistory } from 'vue-router'
import AuthView from '@/views/Auth.vue'
import HomeView from '@/views/Home.vue'
import NavigationView from '@/views/Navigation.vue'
import UserAddView from '@/views/user/Add.vue'
import UserDeleteView from '@/views/user/Delete.vue'
import UserUpdateView from '@/views/user/Update.vue'
import DocumentListView from '@/views/document/List.vue'
import DocumentDeleteView from '@/views/document/Delete.vue'
import DocumentPageView from '@/views/document/Page.vue'
import ErrorPageView from '@/views/Error.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: HomeView,
            children: [{
                path: '/navigation',
                name: 'navigation',
                component: NavigationView
            }, {
                path: '/user/add',
                name: 'user_add',
                component: UserAddView,
            }, {
                path: '/user/delete',
                name: 'user_delete',
                component: UserDeleteView
            }, {
                path: '/user/update',
                name: 'user_update',
                component: UserUpdateView
            }, {
                path: '/document/list',
                name: 'document_list',
                component: DocumentListView
            }, {
                path: '/document/delete',
                name: 'document_delete',
                component: DocumentDeleteView
            }]
        },
        {
            path: '/auth',
            name: 'auth',
            component: AuthView
        },
        {
            path: '/:file_name/:lang',
            name: 'page',
            component: DocumentPageView
        },
        {
            path: '/:pathMatch(.*)*',
            name: '404',
            component: ErrorPageView
        }
    ]
})

export default router
