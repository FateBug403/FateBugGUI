const routers = [
    {
        ID: 1,
        CreatedAt: '2022-12-12T14:39:40.788+08:00',
        UpdatedAt: '2022-12-12T14:39:40.788+08:00',
        parentId: '0',
        path: 'dashboard',
        name: 'dashboard',
        hidden: false,
        component: 'view/dashboard/index.vue',
        sort: 1,
        meta: {
            activeName: '',
            keepAlive: false,
            defaultMenu: false,
            title: '仪表盘',
            icon: 'odometer',
            closeTab: false
        },
        authoritys: null,
        menuBtn: null,
        menuId: '1',
        children: null,
        parameters: [],
        btns: null
    },
    // {
    //     "ID": 2,
    //     "parentId": "0",
    //     "path": "discover",
    //     "name": "discover",
    //     "hidden": false,
    //     "component": "view/discover/discover.vue",
    //     "sort": 5,
    //     "meta": {
    //         "activeName": "",
    //         "keepAlive": true,
    //         "defaultMenu": false,
    //         "title": "资产搜索",
    //         "icon": "tools",
    //         "closeTab": false
    //     },
    //     "authoritys": null,
    //     "menuBtn": null,
    //     "menuId": "2",
    //     "children": null,
    //     "parameters": [],
    //     "btns": null
    // },
    {
        "ID": 3,
        "parentId": "0",
        "path": "afrog",
        "name": "afrog",
        "hidden": false,
        "component": "view/afrog/afrog.vue",
        "sort": 5,
        "meta": {
            "activeName": "",
            "keepAlive": true,
            "defaultMenu": false,
            "title": "POC管理",
            "icon": "tools",
            "closeTab": false
        },
        "authoritys": null,
        "menuBtn": null,
        "menuId": "3",
        "children": null,
        "parameters": [],
        "btns": null
    },
]

export default routers
