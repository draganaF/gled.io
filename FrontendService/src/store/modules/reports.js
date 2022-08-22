import axios from "axios";
import { RECENSIONS_SERVICE_URL } from "../../url";

const state = {
  report: null,
  reports: null,
  result: null
};

const getters = {
  getReport: state => state.report,
  getReports: state => state.reports,
  getResult: state => state.result
};

const actions = {

  fetchReports: (context) => {
    axios.get(`${RECENSIONS_SERVICE_URL}/reports/all`)
      .then(response => {
        context.commit('setReports', response.data);
      })
      .catch(error => {
        context.commit('setResult', {
          label: 'fetch',
          ok: false,
          message: error.response.data
        });
      })
  },

  createReport: (context, report) => {
    axios.post(`${RECENSIONS_SERVICE_URL}/reports/create`, report)
      .then(response => {
       
        context.commit('setReport', response.data);
        context.commit('setResult', {
          label: 'create',
          ok: !response.data.Err ? true : false,
          message: !response.data.Err ?
            'You have successfully reported that recension.' :
            `You have already reported this recension.`
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

  deleteReport: (context, reportId) => {
    axios.delete(`${RECENSIONS_SERVICE_URL}/reports/delete/${reportId}`)
      .then(response => {
       
        context.commit('setReport', response.data);
        context.commit('setResult', {
          label: 'delete',
          ok: !response.data.Err ? true : false,
          message: !response.data.Err ?
            'You have successfully deleted report.' :
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
  },

  silentDeleteReport: (context, reportId) => {
    axios.delete(`${RECENSIONS_SERVICE_URL}/reports/delete/${reportId}`)
      .then(response => {
        
        context.commit('setResult', {
          label: 'silent_delete',
          ok: !response.data.Err ? true : false,
          message: ``
        });
      })
      .catch(error => {
      })
  }
}

const mutations = {
  setReport: (state, response) => {
    state.report = response;
  },

  setReports: (state, response) => {
    state.reports = response;
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