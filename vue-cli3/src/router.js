import Vue from 'vue'
import VueRouter from 'vue-router'
// ルート用のコンポーネントを読み込む
import Home from '@/pages/Home'
import AllNote from '@/pages/Note/AllNote'
import NewNote from '@/pages/Note/NewNote'
import NoteBook from '@/pages/Note/NoteBook'
import Tag from '@/pages/Note/Tag'
import Game from '@/pages/Game'
import Tool from '@/pages/Tool'
// 最初にプラグインとして登録
Vue.use(VueRouter)
// VueRouterインスタンスを生成する
const router = new VueRouter({
    // URLのパスと紐づくコンポーネントをマッピング
    routes: [
        {
            path: '/',
            component: Home
        },
        {
            path: '/note/all',
            component: AllNote,
            // meta: { requireAuth: true},
            children: [
                {
                    path: '/note/new',
                    component: NewNote
                },
                {
                    path: '/note/book',
                    component: NoteBook
                },
                {
                    path: '/note/tag',
                    component: Tag
                },
            ]
        },
        {
            path: '/game',
            component: Game
        },
        {
            path: '/tool',
            component: Tool
        }
    ]
})
// 生成したVueRouterインスタンスをエクスポート
export default router