import axios from "axios";

const state = {
  projection: null,
  projections: null,
  result: null
};

const getters = {
  getProjection: state => state.projection,
  getProjections: state => state.projections,
  getResult: state => state.result
};

const actions = {
  fetchProjections: (context) => {
    axios.get(`http://localhost:8082/api/projections`)
    .then(response => {
        console.log(response)
        context.commit('setProjections', response.data);
    })
    .catch(error => {
        context.commit('setResult', { label: 'fetch', ok: false });
    });
  },

  fetchProjectionById: (context, id) => {
    axios.get(`http://localhost:8082/api/projections/${id}`)
    .then(response => {
        context.commit('setProjection', response.data);
    })
    .catch(error => {
        context.commit('setResult', { label: 'fetch', ok: false });
    });
  },

  createProjection: (context, projection) => {
    axios.post(`http://localhost:8082/api/projections`, projection)
    .then(response => {
      context.commit('setResult', { 
        label: 'create',
        ok: true,
        message: "You have successfully added new projection"
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
  deleteProjection: (context, id) => {
    axios.delete(`http://localhost:8082/api/projections/` + id)
    .then(response => {
      context.commit('setResult', { 
        label: 'delete',
        ok: true,
        message: "You have successfully deleted projection"
      });
    })
    .catch(error => {
      context.commit('setResult', { 
        label: 'delete', 
        ok: false,
        message: error.response.data  
      });
    });
  }
}


const mutations = {

  setProjections: (state, response) => {
    state.projections = response;
  },

  setProjection: (state, response) => {
    state.projection = response;
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