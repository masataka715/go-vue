import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        is_auth: false,
        show_login_form: false,
        go_domain: "http://localhost:5000",
        test_auth_id: 9185491
    },
    getters: {
        checkAuth: state => {
            const session_id = window.sessionStorage.getItem(["user_id"]);
            if (session_id) {
                state.is_auth = true
            } else {
                state.is_auth = false
            }
            return state.is_auth
        },
    },
    mutations: {
        login(state) {
            state.is_auth = true
        },
        logout(state) {
            state.is_auth = false
        },
        showLoginForm(state) {
            state.show_login_form = true
        },
        hideLoginForm(state) {
            state.show_login_form = false
        }
    },
})
// 生成したVueRouterインスタンスをエクスポート
export default store