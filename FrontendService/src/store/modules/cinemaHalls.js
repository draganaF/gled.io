import axios from "axios";
import { PROJECTION_SERVICE_URL } from "../../url";
const state = {
  hall: null,
  halls: null,
  result: null
};

const getters = {
  getHall: state => state.hall,
  getHalls: state => state.halls,
  getResult: state => state.result
};

const actions = {
  fetchHalls: (context) => {
    axios.get(`${PROJECTION_SERVICE_URL}/cinema-halls`)
    .then(response => {
        console.log(response)
        context.commit('setHalls', response.data);
    })
    .catch(error => {
        context.commit('setResult', { label: 'fetch', ok: false });
    });
  },

  fetchHallById: (context, id) => {
    axios.get(`${PROJECTION_SERVICE_URL}/cinema-halls/${id}`)
    .then(response => {
        context.commit('setHall', response.data);
    })
    .catch(error => {
        context.commit('setResult', { label: 'fetch', ok: false });
    });
  },
}


const mutations = {

  setHalls: (state, response) => {
    state.halls = response;
  },

  setHall: (state, response) => {
    state.hall = response;
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