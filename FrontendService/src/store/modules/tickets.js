import axios from "axios";
import { TICKETS_SERVICE_URL } from "../../url";

const state = {
  ticket: null,
  tickets: null,
  result: null
};

const getters = {
  getTicket: state => state.ticket,
  getTickets: state => state.tickets,
  getResult: state => state.result
};

const actions = {
  fetchTickets: (context) => {
    axios.get(`${TICKETS_SERVICE_URL}/tickets`)
    .then(response => {
        context.commit('setTickets', response.data);
    })
    .catch(error => {
        context.commit('setResult', { label: 'fetch', ok: false });
    });
  },
  fetchUsersTickets: (context, id ) => {
    axios.get(`${TICKETS_SERVICE_URL}/tickets/users/${id}`)
    .then(response => {
        context.commit('setTickets', response.data);
    })
    .catch(error => {
        context.commit('setResult', { label: 'fetch', ok: false });
    });
  },

  fetchTicketById: (context, id) => {
    axios.get(`${TICKETS_SERVICE_URL}/tickets/${id}`)
    .then(response => {
        context.commit('setTicket', response.data);
    })
    .catch(error => {
        context.commit('setResult', { label: 'fetch', ok: false });
    });
  },

  fetchTicketsByProjectionId: (context, projectionId) => {
    axios.get(`${TICKETS_SERVICE_URL}/tickets/projections/${projectionId}`)
      .then(response => {
        context.commit('setTickets', response.data);
      })
      .catch(error => {
        context.commit('setResult', { label: 'fetch', ok: false });
      })
  },

  createTicket: (context, ticket) => {
    axios.post(`${TICKETS_SERVICE_URL}/tickets`, ticket)
    .then(response => {
        context.commit('setResult', {
          label : 'create',
          ok: true,
          message: "You have successfully added ticket"
        });
    })
    .catch(error => {
        console.log(error)
        context.commit('setResult', { 
          label: 'create',
          ok: false,
          message: error.response.data
         });
    });
  }, 
  buyTicket: (context, id) => {
    axios.get(`${TICKETS_SERVICE_URL}/tickets/buy-reserved-ticket/${id}`)
    .then(response => {
        context.commit('setResult', {label: 'buy', ok: true, message: "You have succesfuly bought ticket"});
    })
    .catch(error => {
        context.commit('setResult', { label: 'buy', ok: false });
    });
  },

}


const mutations = {

  setTickets: (state, response) => {
    state.tickets = response;
  },

  setTicket: (state, response) => {
    state.ticket = response;
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