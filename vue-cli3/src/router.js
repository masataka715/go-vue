import Vue from 'vue'
import VueRouter from 'vue-router'
// ルート用のコンポーネントを読み込む
import Home from '@/views/Home'
import Game from '@/views/Game'
import Tool from '@/views/Tool'
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