import { defineStore } from 'pinia'

type MenuItem = {
    title: string,
    items: string[]
}

type NotificationItem = {
    type: 'error' | 'success' | 'info' | 'warning',
    title: string,
    message: string
}

type MessageBoxItem = {
    type: 'error' | 'success' | 'info' | 'warning',
    message: string,
    title: string,
    confirmButtonText: string
}

export const langStore = defineStore('menu', () => {
    const menuItems = (lang: string = 'en') => {
        let result: {
            sideMenu: {
                collapseButton: {
                    expand: string,
                    fold: string
                },
                user: MenuItem,
                document: MenuItem,
                setting: MenuItem
            },
            topMenu: {
                title: string,
                logout: string
            }
        } = {
            sideMenu: {
                collapseButton: {
                    expand: 'Expand',
                    fold: 'Fold',
                },
                user: {
                    title: 'User',
                    items: ['Add', 'Delete', 'Update']
                },
                document: {
                    title: 'Document',
                    items: ['List', 'Delete']
                },
                setting: {
                    title: 'Setting',
                    items: ['Light/Dark', 'Language']
                }
            },
            topMenu: {
                title: 'Management',
                logout: 'Logout'
            }
        }
        if (lang === 'zh') {
            result = {
                sideMenu: {
                    collapseButton: {
                        expand: '收起',
                        fold: '展开',
                    },
                    user: {
                        title: '用户',
                        items: ['添加', '删除', '更新']
                    },
                    document: {
                        title: '文档',
                        items: ['列表', '删除']
                    },
                    setting: {
                        title: '设置',
                        items: ['亮/暗', '语言']
                    }
                },
                topMenu: {
                    title: '管理',
                    logout: '注销'
                }
            }
        }
        return result
    }
    const notificationItems = (lang: string = 'en') => {
        let result: {
            [key in string]: {
                success?: NotificationItem,
                error?: NotificationItem,
                info?: NotificationItem,
                warning?: NotificationItem
            }
        } = {
            ping: {
                success: {
                    type: 'success',
                    title: 'Success',
                    message: 'The backend connection is successful'
                },
                error: {
                    type: 'error',
                    title: 'Error',
                    message: 'Backend connection failed'
                }
            },
            login: {
                success: {
                    type: 'success',
                    title: 'Success',
                    message: 'Logined!'
                },
                error: {
                    type: 'error',
                    title: 'Error',
                    message: 'Login failed'
                }
            },
            auth: {
                error: {
                    type: 'error',
                    message: 'The token authentication is incorrect, please log in again',
                    title: 'Error'
                }
            },
            deleteUser: {
                error: {
                    type: 'error',
                    message: 'Failed to delete user',
                    title: 'Error'
                },
                success: {
                    type: 'success',
                    message: 'The user was deleted successfully',
                    title: 'Success'
                }
            },
            deleteAdmin: {
                error: {
                    type: 'error',
                    message: 'Deleting the admin user is not allowed',
                    title: 'Error'
                }
            },
            deleteSelf: {
                error: {
                    type: 'error',
                    message: 'Deleting the current user is not allowed',
                    title: 'Error'
                }
            },
            updateUser: {
                error: {
                    type: 'error',
                    message: 'Failed to update user information',
                    title: 'Error'
                },
                success: {
                    type: 'success',
                    message: 'Update user information successful',
                    title: 'Success'
                }
            },
            unknownPermission: {
                error: {
                    type: 'error',
                    message: 'Permissions other than admin and custom are not supported',
                    title: 'Error'
                }
            },
            notAllowedToChangeAdminInfo: {
                error: {
                    type: 'error',
                    message: 'It is not allowed to change the information of the admin user',
                    title: 'Error'
                }
            },
            notAllowedChangePermission: {
                error: {
                    type: 'error',
                    message: 'Changing the permissions of the current user is not allowed',
                    title: 'Error'
                }
            },
            newUserIsExists: {
                error: {
                    type: 'error',
                    message: 'A new user already exists',
                    title: 'error'
                }
            },
            createdUser: {
                success: {
                    type: 'success',
                    message: 'The user is created successfully',
                    title: 'Success'
                }
            },
            createUserFailed: {
                error: {
                    type: 'error',
                    message: 'User creation failed',
                    title: 'Error'
                }
            },
            MisssingNameFiled: {
                error: {
                    type: 'error',
                    message: 'The user name cannot be empty',
                    title: 'Error'
                }
            },
            deleteDocumentFailed: {
                error: {
                    type: 'error',
                    message: 'Failed to delete document',
                    title: 'Error'
                }
            },
            deleteDocumentSuccess: {
                success: {
                    type: 'success',
                    message: 'The document was deleted successfully',
                    title: 'Success'
                }
            }
        }
        if (lang === 'zh') {
            result = {
                ping: {
                    success: {
                        type: 'success',
                        title: '成功',
                        message: '连接后端成功'
                    },
                    error: {
                        type: 'error',
                        title: '失败',
                        message: '连接后端失败'
                    }
                },
                login: {
                    success: {
                        type: 'success',
                        title: '成功',
                        message: '登录成功'
                    },
                    error: {
                        type: 'error',
                        title: '失败',
                        message: '登录失败'
                    }
                },
                auth: {
                    error: {
                        type: 'error',
                        message: '用户鉴权失败，请重新登录',
                        title: '鉴权失败'
                    }
                },
                deleteUser: {
                    error: {
                        type: 'error',
                        message: '删除用户失败',
                        title: '失败'
                    },
                    success: {
                        type: 'success',
                        message: '删除用户成功',
                        title: '成功'
                    }
                },
                deleteAdmin: {
                    error: {
                        type: 'error',
                        message: 'admin用户不允许删除',
                        title: '失败'
                    }
                },
                deleteSelf: {
                    error: {
                        type: 'error',
                        message: '不允许删除当前用户',
                        title: '失败'
                    }
                },
                updateUser: {
                    error: {
                        type: 'error',
                        message: '更新用户信息失败',
                        title: '失败'
                    },
                    success: {
                        type: 'success',
                        message: '更新用户信息成功',
                        title: '成功'
                    }
                },
                unknownPermission: {
                    error: {
                        type: 'error',
                        message: '不支持该用户权限，只支持admin和custom',
                        title: '失败'
                    }
                },
                notAllowedToChangAdminInfo: {
                    error: {
                        type: 'error',
                        message: '不允许更改admin用户的信息',
                        title: '失败'
                    }
                },
                notAllowedChangePermission: {
                    error: {
                        type: 'error',
                        message: '当前用户不允许改变权限',
                        title: '失败'
                    }
                },
                newUserIsExists: {
                    error: {
                        type: 'error',
                        message: '新建用户已存在!',
                        title: '失败'
                    }
                },
                createdUser: {
                    success: {
                        type: 'success',
                        message: '成功新建用户',
                        title: '成功'
                    }
                },
                createUserFailed: {
                    error: {
                        type: 'error',
                        message: '创建用户失败',
                        title: '失败'
                    }
                },
                MisssingNameFiled: {
                    error: {
                        type: 'error',
                        message: '用户名不能为空',
                        title: '失败'
                    }
                },
                deleteDocumentFailed: {
                    error: {
                        type: 'error',
                        message: '文档删除失败',
                        title: '失败'
                    }
                },
                deleteDocumentSuccess: {
                    success: {
                        type: 'success',
                        message: '文档删除成功',
                        title: '成功'
                    }
                }
            }
        }
        return result
    }
    const loginFormItems = (lang: string = 'en') => {
        let result: {
            header: string,
            name_placeholder: string,
            password_placeholder: string,
            button: string
        } = {
            header: 'Login',
            name_placeholder: 'User name',
            password_placeholder: 'Password',
            button: 'Login'
        }
        if (lang === 'zh') {
            result = {
                header: '登录',
                name_placeholder: '用户名',
                password_placeholder: '密码',
                button: '登录'
            }
        }
        return result
    }
    const navigationItems = (lang: string = 'en') => {
        let result: {
            title: string,
            main_content: string[]
        } = {
            title: 'Online clipboard',
            main_content: ['Online clipboard service management interface', 'User management and article management']
        }
        if (lang === 'zh') {
            result = {
                title: '在线剪切板',
                main_content: ['在线剪切板服务管理界面', '用户管理以及文章管理']
            }
        }
        return result
    }
    const tableItems = (lang: string = 'en') => {
        let result: {
            user: {
                [key: string]: {
                    id: string,
                    name: string,
                    permission: string,
                    operate: string,
                    search?: string,
                    button: {
                        [key: string]: string
                    },
                    edit?: {
                        title: string,
                        name: string,
                        password: string,
                        permission: string,
                        button: {
                            cancel: string,
                            confirm: string
                        },
                        placeHolder: {
                            name: string,
                            password: string
                        }
                    },
                    add?: {
                        title: string,
                        name: string,
                        password: string,
                        permission: string,
                        button: {
                            cancel: string,
                            confirm: string
                        },
                        placeHolder: {
                            name: string,
                            password: string
                        }
                    }
                }
            },
            document: {
                [key: string]: {
                    id: string,
                    path: string,
                    upload_user: string,
                    operate?: string,
                    search?: string,
                    delete?: string,
                    button?: {
                        [key: string]: string
                    }
                }
            }
        } = {
            user: {
                delete: {
                    id: 'ID',
                    name: 'Name',
                    permission: 'Permission',
                    operate: 'Operate',
                    search: 'Type keyword to search',
                    button: {
                        delete: 'Delete'
                    }
                },
                update: {
                    id: 'ID',
                    name: 'Name',
                    permission: 'Permission',
                    operate: 'Operate',
                    search: 'Type keyword to search',
                    button: {
                        edit: 'Edit'
                    },
                    edit: {
                        title: 'Edit User',
                        name: 'Name',
                        password: 'Password',
                        permission: 'Permission',
                        button: {
                            cancel: 'Cancel',
                            confirm: 'Submit'
                        },
                        placeHolder: {
                            name: 'User name',
                            password: 'The default is the original password.(empty value)'
                        }
                    }
                },
                add: {
                    id: 'ID',
                    name: 'Name',
                    permission: 'Permission',
                    operate: 'Operate',
                    search: 'Type keyword to search',
                    button: {
                        add: 'Add user'
                    },
                    add: {
                        title: 'Add a new user',
                        name: 'Name',
                        password: 'Password',
                        permission: 'Permission',
                        button: {
                            cancel: 'Cancel',
                            confirm: 'Submit'
                        },
                        placeHolder: {
                            name: 'User name',
                            password: 'The default is the original password.(empty value)'
                        }
                    }
                }
            },
            document: {
                list: {
                    id: 'ID',
                    path: 'Path',
                    upload_user: 'Upload user',
                    search: 'Type keyword to search',
                },
                delete: {
                    id: 'ID',
                    path: 'Path',
                    upload_user: 'Upload user',
                    operate: 'Operate',
                    delete: 'Delete',
                    search: 'Type keyword to search',
                }
            }
        }
        if (lang === 'zh') {
            result = {
                user: {
                    delete: {
                        id: 'ID',
                        name: '用户名',
                        permission: '权限',
                        operate: '操作',
                        search: '输入关键字查询',
                        button: {
                            delete: '删除'
                        }
                    },
                    update: {
                        id: 'ID',
                        name: '用户名',
                        permission: '权限',
                        operate: '操作',
                        search: '输入关键字查询',
                        button: {
                            edit: '编辑'
                        },
                        edit: {
                            title: '编辑用户',
                            name: '用户名',
                            password: '密码',
                            permission: '权限',
                            button: {
                                cancel: '取消',
                                confirm: '提交'
                            },
                            placeHolder: {
                                name: '用户名',
                                password: '默认为当前密码。(空值)'
                            },
                        }
                    },
                    add: {
                        id: 'ID',
                        name: '用户名',
                        permission: '权限',
                        operate: '操作',
                        search: '输入关键字查询',
                        button: {
                            add: '添加用户'
                        },
                        add: {
                            title: '新增用户',
                            name: '用户名',
                            password: '密码',
                            permission: '权限',
                            button: {
                                cancel: '取消',
                                confirm: '提交'
                            },
                            placeHolder: {
                                name: '用户名',
                                password: '默认当前密码。(空值)'
                            }
                        }
                    }
                },
                document: {
                    list: {
                        id: 'ID',
                        path: '路径',
                        upload_user: '上传用户',
                        search: '输入关键字查询',
                    },
                    delete: {
                        id: 'ID',
                        path: '路径',
                        upload_user: '上传用户',
                        operate: '操作',
                        search: '输入关键字查询',
                        delete: '删除'
                    }
                }
            }
        }
        return result
    }
    const messageBoxItems = (lang: string = 'en') => {
        let result: {
            [key: string]: {
                alert?: {
                    error?: MessageBoxItem,
                    success?: MessageBoxItem,
                    info?: MessageBoxItem,
                    warning?: MessageBoxItem
                }
            }
        } = {
            notAllowChangeOtherUser: {
                alert: {
                    warning: {
                        type: 'warning',
                        message: 'The current user is not allowed to change the information of other users',
                        confirmButtonText: 'OK',
                        title: 'Warning'
                    }
                }
            },
            userNameIsEmpty: {
                alert: {
                    warning: {
                        type: 'warning',
                        message: 'The user name cannot be empty',
                        confirmButtonText: 'OK',
                        title: 'Warning'
                    }
                }
            },
            onlyAllowAdminToRegisterNewUser: {
                alert: {
                    warning: {
                        type: 'warning',
                        message: 'Only admin is allowed to register new users',
                        confirmButtonText: 'OK',
                        title: 'Warning'
                    }
                }
            },
            getContentFailed: {
                alert: {
                    error: {
                        type: 'error',
                        message: 'Failed to get content',
                        confirmButtonText: 'OK',
                        title: 'Error'
                    }
                }
            },
            noPermissionToDeletePage: {
                alert: {
                    warning: {
                        type: 'warning',
                        message: 'You do not have permission to delete the document',
                        confirmButtonText: 'OK',
                        title: 'Error'
                    }
                }
            },
            notAllowedChangePermission: {
                alert: {
                    warning: {
                        type: 'warning',
                        message: 'Modification of this user right is not allowed',
                        confirmButtonText: 'OK',
                        title: 'Warning'
                    }
                }
            }
        }
        if (lang === 'zh') {
            result = {
                notAllowChangeOtherUser: {
                    alert: {
                        warning: {
                            type: 'warning',
                            message: '当前用户不允许修改其他用户信息',
                            confirmButtonText: '好的',
                            title: '警告'
                        }
                    }
                },
                userNameIsEmpty: {
                    alert: {
                        warning: {
                            type: 'warning',
                            message: '用户名不能为空',
                            confirmButtonText: '好的',
                            title: '警告'
                        }
                    }
                },
                onlyAllowAdminToRegisterNewUser: {
                    alert: {
                        warning: {
                            type: 'warning',
                            message: '仅允许admin用户注册新用户',
                            confirmButtonText: '好的',
                            title: '警告'
                        }
                    }
                },
                getContentFailed: {
                    alert: {
                        error: {
                            type: 'error',
                            message: '获取内容失败',
                            confirmButtonText: '好的',
                            title: '失败'
                        }
                    }
                },
                noPermissionToDeletePage: {
                    alert: {
                        warning: {
                            type: 'warning',
                            message: '没有删除该文档的权限',
                            confirmButtonText: '好的',
                            title: '警告'
                        }
                    }
                },
                notAllowedChangePermission: {
                    alert: {
                        warning: {
                            type: 'warning',
                            message: '不允许修改该用户权限',
                            confirmButtonText: '好的',
                            title: '警告'
                        }
                    }
                }
            }
        }
        return result
    }
    const errorPageItems = (lang: string = 'en') => {
        let result: {
            title: string,
            sub_title: string,
            button: string
        } = {
            title: '404 Not Found',
            sub_title: 'Page is not found',
            button: 'Home'
        }
        if (lang === 'zh') {
            result = {
                title: '404未找到页面',
                sub_title: '页面未找到',
                button: '首页'
            }
        }
        return result
    }
    return { menuItems, notificationItems, loginFormItems, navigationItems, tableItems, messageBoxItems, errorPageItems }
})