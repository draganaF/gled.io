import Vue from 'vue'
import Vuex from 'vuex'
import authentication from './modules/authentication'
import projections from './modules/projections'
import movies from './modules/movies'
import cinemaHalls from './modules/cinemaHalls'
import users from './modules/users'
import tickets from './modules/tickets'
import recensions from './modules/recensions'
import reports from './modules/reports'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    authentication,
    projections,
    cinemaHalls,
    movies,
    users,
    tickets,
    recensions,
    reports
  }
});
