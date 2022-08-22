import axios from "axios";
import { PROJECTION_SERVICE_URL } from "../../url";
const state = {
  movie: null,
  movies: null,
  result: null
};

const getters = {
  getMovie: state => state.movie,
  getMovies: state => state.movies,
  getResult: state => state.result
};

const actions = {
  fetchMovies: (context) => {
    axios.get(`${PROJECTION_SERVICE_URL}/movies`)
    .then(response => {
        context.commit('setMovies', response.data);
    })
    .catch(error => {
        context.commit('setResult', { label: 'fetch', ok: false });
    });
  },

  fetchMovieById: (context, id) => {
    axios.get(`${PROJECTION_SERVICE_URL}/movies/${id}`)
    .then(response => {
        context.commit('setMovie', response.data);
    })
    .catch(error => {
        context.commit('setResult', { label: 'fetch', ok: false });
    });
  },

  createMovie: (context, movie) => {
    axios.post(`${PROJECTION_SERVICE_URL}/movies`, movie)
    .then(response => {
        context.commit('setResult', {
          label : 'create',
          ok: true,
          message: "You have successfully added new movie"
        });
    })
    .catch(error => {
      
        context.commit('setResult', { 
          label: 'create',
          ok: false,
          message: error.response.data
         });
    });
  }, 

  updateMovie: (context, movie) => {
    axios.put(`${PROJECTION_SERVICE_URL}/movies`, movie)
    .then(response => {
        context.commit('setResult', {
          label : 'update',
          ok: true,
          message: "You have successfully updated movie"
        });
    })
    .catch(error => {
      
        context.commit('setResult', { 
          label: 'update',
          ok: false,
          message: error.response.data
         });
    });
  }, 
}


const mutations = {

  setMovies: (state, response) => {
    state.movies = response;
  },

  setMovie: (state, response) => {
    state.movie = response;
  },

  setResult: (state, response) => {
    state.result = response;
  }
};

export default {
  state: state,
  getters: getters,
  actions: actions,
  mutations: mutations,
  namespaced: true
};