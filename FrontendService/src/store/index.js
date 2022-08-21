import Vue from 'vue'
import Vuex from 'vuex'
import authentication from './modules/authentication'
import projections from './modules/projections'
import movies from './modules/movies'
import cinemaHalls from './modules/cinemaHalls'
Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    authentication,
    projections,
    cinemaHalls,
    movies
  }
});
