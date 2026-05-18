import './style.css';
import {GetSSHHosts, ConnectToHost} from '../wailsjs/go/main/App';

const app = document.querySelector('#app');

let allHosts = [];

async function init() {
  app.innerHTML = `
    <div class="header">
      <h1>SSH Connect</h1>
      <p>Select a host to launch a terminal session</p>
    </div>
    <div class="search-container">
      <input type="text" class="search-input" id="search" placeholder="Search hosts..." autocomplete="off">
    </div>
    <div class="hosts-list" id="hosts-list">
      <div class="empty-state">Loading hosts...</div>
    </div>
  `;

  const searchInput = document.getElementById('search');
  const hostsList = document.getElementById('hosts-list');

  if (searchInput) {
    searchInput.addEventListener('input', (e) => {
      const query = e.target.value.toLowerCase();
      const filtered = allHosts.filter(h => h.toLowerCase().includes(query));
      renderHosts(filtered);
    });
  }

  try {
    allHosts = await GetSSHHosts();
    renderHosts(allHosts);
  } catch (err) {
    if (hostsList) {
      hostsList.innerHTML = `<div class="empty-state">Error: ${err}</div>`;
    }
  }
}

function renderHosts(hosts) {
  const hostsList = document.getElementById('hosts-list');
  if (!hostsList) return;

  if (hosts.length === 0) {
    hostsList.innerHTML = `<div class="empty-state">No hosts found</div>`;
    return;
  }

  hostsList.innerHTML = hosts.map((host, index) => `
    <div class="host-card" style="animation-delay: ${index * 0.05}s" onclick="connect('${host}')">
      <div class="host-icon">🖥️</div>
      <div class="host-name">${host}</div>
    </div>
  `).join('');
}

window.connect = async (host) => {
  try {
    await ConnectToHost(host);
  } catch (err) {
    console.error(err);
  }
};

init();
