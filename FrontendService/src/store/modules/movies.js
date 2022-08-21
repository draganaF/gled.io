import axios from "axios";

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
    axios.get(`http://localhost:8082/api/movies`)
    .then(response => {
        console.log(response)
        context.commit('setMovies', response.data);
    })
    .catch(error => {
        context.commit('setResult', { label: 'fetch', ok: false });
    });
  },

  fetchMovieById: (context, id) => {
    axios.get(`http://localhost:8082/api/movies/${id}`)
    .then(response => {
        context.commit('setMovie', response.data);
    })
    .catch(error => {
        context.commit('setResult', { label: 'fetch', ok: false });
    });
  },

  createMovie: (context, movie) => {
    axios.post(`http://localhost:8082/api/movies`, movie)
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
  }
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