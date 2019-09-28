import Vue from 'vue'
import VueRouter from 'vue-router'
// ルート用のコンポーネントを読み込む
import Home from '@/pages/Home'
import NoteApp from '@/pages/Note/NoteApp'
import AllNote from '@/pages/Note/AllNote'
import NoteDetail from '@/pages/Note/NoteDetail'
import NewNote from '@/pages/Note/NewNote'
import NoteBook from '@/pages/Note/NoteBook'
import NoteTag from '@/pages/Note/NoteTag'
import NoteGarbage from '@/pages/Note/NoteGarbage'
import GarbageDetail from '@/pages/Note/GarbageDetail'
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
            path: '/note',
            redirect: '/note/all',
            component: NoteApp,
            // meta: { requireAuth: true},
            children: [
                {
                    name: 'note_all',
                    path: '/note/all',
                    component: AllNote,
                    children: [
                        {
                            name: 'note_details',
                            path: '/note/all/:id',
                            component: NoteDetail
                        }
                    ]
                },
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
                    component: NoteTag
                },
                {
                    path: '/note/garbage',
                    component: NoteGarbage,
                    children: [
                        {
                            name: 'garbage_details',
                            path: '/note/garbage/:id',
                            component: GarbageDetail
                        }
                    ]
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