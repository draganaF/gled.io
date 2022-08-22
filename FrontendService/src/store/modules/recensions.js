import axios from "axios";
import { RECENSIONS_SERVICE_URL } from "../../url";

const state = {
  score: 0,
  recension: null,
  recensions: null,
  result: null
};

const getters = {
  getScore: state => state.score,
  getRecension: state => state.recension,
  getRecensions: state => state.recensions,
  getResult: state => state.result
};

const actions = {
  fetchScoreByMovieId: (context, movieId) => {
    axios.get(`${RECENSIONS_SERVICE_URL}/recensions/score-by-movie-id/${movieId}`)
    .then(response => {
      context.commit('setScore', response.data.Score);
    })
    .catch(error => {
      context.commit('setResult', {
        label: 'score',
        ok: false,
        message: error.response.data
      });
    })
  },
  fetchByMovie: (context, movieId) => {
    axios.get(`${RECENSIONS_SERVICE_URL}/recensions/by-movie-id/${movieId}`)
      .then(response => {
        console.log(response);
        context.commit('setRecensions', response.data);
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'fetch',
          ok: false,
          message: error.response.data
        });
      })
  },

  fetchRecensions: (context) => {
    axios.get(`${RECENSIONS_SERVICE_URL}/recensions/all`)
      .then(response => {
        console.log(response);
        context.commit('setRecensions', response.data);
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'fetch',
          ok: false,
          message: error.response.data
        });
      })
  },

  createRecension: (context, recension) => {
    axios.post(`${RECENSIONS_SERVICE_URL}/recensions/create`, recension)
      .then(response => {
        console.log(response);
        context.commit('setRecension', response.data);
        context.commit('setResult', {
          label: 'create',
          ok: !response.data.Err ? true : false,
          message: !response.data.Err ?
            'You have successfully gave a new recension.' :
            `You have already gave recension for this movie.`
        });

      })
      .catch(error => {
        context.commit('setResult', {
          label: 'create',
          ok: false,
          message: error.response.data
        });
      })
  },

  deleteRecension: (context, recensionId) => {
    axios.delete(`${RECENSIONS_SERVICE_URL}/recensions/delete/${recensionId}`)
      .then(response => {
        console.log(response);
        context.commit('setRecension', response.data);
        context.commit('setResult', {
          label: 'delete',
          ok: !response.data.Err ? true : false,
          message: !response.data.Err ?
            'You have successfully deleted recension.' :
            response.data.Err
        });

      })
      .catch(error => {
        context.commit('setResult', {
          label: 'delete',
          ok: false,
          message: error.response.data
        });
      })
  }

};

const mutations = {
  setScore: (state, response) => {
    state.score = response;
  },

  setRecension: (state, response) => {
    state.recension = response;
  },

  setRecensions: (state, response) => {
    state.recensions = response;
  },

  setResult: (state, response) => {
    state.result = response;
  }
}

export default {
  state: state,
  getters: getters,
  actions: actions,
  mutations: mutations,
  namespaced: true
};