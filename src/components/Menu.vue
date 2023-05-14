<script setup lang="ts">
import { useDark, useToggle } from '@vueuse/core'
import {
    Expand,
    Fold,
    Plus,
    Delete,
    Edit,
    User,
    Document,
    Setting,
    Sunny,
    Moon,
    Switch
} from '@element-plus/icons-vue'
import { langStore } from '@/stores/lang'
import { ref } from 'vue'
import router from '@/router'
import { RouterView, RouterLink } from 'vue-router'

// 响应式数据
const page = ref('navigation')
const isCollapse = ref(true)
const _userName = localStorage.getItem('user')
const userName = ref(_userName === undefined ? 'unknown' : _userName)
const lang = ref(localStorage.getItem('lang') == null ? 'en' : localStorage.getItem('lang') as string)
const dark = ref(useDark() as unknown as boolean)

// 获取界面文字
const items = langStore().menuItems(lang.value)

// 展开和收缩侧边栏
const handleChange = () => {
    isCollapse.value = !isCollapse.value
}

// 判断是否是暗色
const isDark = useDark({
    onChanged: (_dark: boolean) => {
        dark.value = _dark
    }
})
// 改变颜色
const toggleDark = useToggle(isDark)

// 改变语言
const toggleLang = () => {
    lang.value = (lang.value == 'en' ? 'zh' : 'en') as unknown as string
    localStorage.setItem('lang', lang.value as unknown as string)
    location.reload()
}

// 登出
const logoutHandle = () => {
    localStorage.removeItem('token')
    localStorage.removeItem('role')
    localStorage.removeItem('user')
    router.push('/auth')
}

// 返回主页
const goHome = () => {
    router.push('/')
}

// 改变页面路由
const changePage = (_page: string) => {
    page.value = _page
}
</script>

<template>
    <div class="container">
        <el-header class="topmenu">
            <el-menu mode="horizontal" :ellipsis="false">
                <RouterLink :to="{ name: 'navigation' }">
                    <el-menu-item index="0" @click="goHome">{{ items.topMenu.title }}</el-menu-item>
                </RouterLink>
                <el-menu-item @click="handleChange" index="1">
                    <el-icon v-if="isCollapse">
                        <Expand />
                    </el-icon>
                    <el-icon v-else>
                        <Fold />
                    </el-icon>
                </el-menu-item>
                <div class="flex-grow" />
                <el-sub-menu index="2">
                    <template #title>
                        {{ userName == undefined ? "unknown" : userName }}
                    </template>
                    <el-menu-item index="2-1" @click="logoutHandle">{{ items.topMenu.logout }}</el-menu-item>
                </el-sub-menu>
            </el-menu>
        </el-header>
        <el-container>
            <el-side>
                <el-menu class="sidemenu" :collapse="isCollapse">
                    <!-- Expand button -->
                    <!-- User management -->
                    <el-sub-menu index="0">
                        <template #title>
                            <el-icon>
                                <User />
                            </el-icon>
                            <span>{{ items.sideMenu.user.title }}</span>
                        </template>
                        <RouterLink :to="{ name: 'user_add' }">
                            <el-menu-item index="0-1" @click="changePage('user_add')">
                                <el-icon>
                                    <Plus />
                                </el-icon>
                                <span>{{ items.sideMenu.user.items[0] }}</span>
                            </el-menu-item>
                        </RouterLink>
                        <RouterLink :to="{ name: 'user_delete' }">
                            <el-menu-item index="0-3">
                                <el-icon>
                                    <Delete />
                                </el-icon>
                                <span>{{ items.sideMenu.user.items[1] }}</span>
                            </el-menu-item>
                        </RouterLink>
                        <RouterLink :to="{ name: 'user_update' }">
                            <el-menu-item index="0-4">
                                <el-icon>
                                    <Edit />
                                </el-icon>
                                <span>{{ items.sideMenu.user.items[2] }}</span>
                            </el-menu-item>
                        </RouterLink>
                    </el-sub-menu>
                    <!-- Documents management -->
                    <el-sub-menu index="1">
                        <template #title>
                            <el-icon>
                                <Document />
                            </el-icon>
                            <span>{{ items.sideMenu.document.title }}</span>
                        </template>
                        <RouterLink :to="{ name: 'document_list' }">
                            <el-menu-item index="1-1">
                                <el-icon>
                                    <Document />
                                </el-icon>
                                <span>{{ items.sideMenu.document.items[0] }}</span>
                            </el-menu-item>
                        </RouterLink>
                        <RouterLink :to="{ name: 'document_delete' }">
                            <el-menu-item index="1-2">
                                <el-icon>
                                    <Delete />
                                </el-icon>
                                <span>{{ items.sideMenu.document.items[1] }}</span>
                            </el-menu-item>
                        </RouterLink>
                    </el-sub-menu>
                    <!-- Settings -->
                    <el-sub-menu index="2">
                        <template #title>
                            <el-icon>
                                <Setting />
                            </el-icon>
                            <span>{{ items.sideMenu.setting.title }}</span>
                        </template>
                        <el-menu-item @click="toggleDark()" index="2-1">
                            <el-icon v-if="dark">
                                <Moon />
                            </el-icon>
                            <el-icon v-else>
                                <Sunny />
                            </el-icon>
                            <span>{{ items.sideMenu.setting.items[0] }}</span>
                        </el-menu-item>
                        <el-menu-item @click="toggleLang" index="2-2">
                            <el-icon>
                                <Switch />
                            </el-icon>
                            <span>{{ items.sideMenu.setting.items[1] }}</span>
                        </el-menu-item>
                    </el-sub-menu>
                </el-menu>
            </el-side>
            <el-main>
                <RouterView class="page" />
            </el-main>
        </el-container>
    </div>
</template>

<style scoped>
.container {
    flex-direction: column;
    display: flex;
    height: 100vh;
}

.sidemenu:not(.el-menu--collapse) {
    width: 12rem;
    height: 100%;
}

.sidemenu {
    height: 100%;
}

.topmenu {
    width: 100%;
    margin: 0;
    padding: 0;
}

.topmenu>* {
    height: 100%;
}

.flex-grow {
    flex-grow: 1;
}

.page {
    height: 100%;
}

a {
    text-decoration: none;
}
</style>
