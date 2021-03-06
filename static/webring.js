const WEBRING_API_URL = 'https://jae.fi/webring/members';

const template = document.createElement('template');

template.innerHTML = `
<style>

.ring {
    display: block;
    float: left;
    max-width: 600px;
    margin: 1rem auto;
}

.webring {
    border: 2px solid #fff;
    padding: 1rem;
    text-align: left;
    font: 100% system-ui, sans-serif;
}

.webring a {
    color: #F4511E;
}

</style>

<div class="ring">
    <div class="webring">
        <div id="copy">


        </div>
    </div>
</div>`;

 class WebRing extends HTMLElement {
    connectedCallback() {
        this.attachShadow({ mode: 'open' });
        this.shadowRoot.appendChild(template.content.cloneNode(true));

        const thisSite = this.getAttribute('site');

        fetch(WEBRING_API_URL)
            .then((response) => response.json())
            .then((sites) => {
                const matchedSiteIndex = sites.members.findIndex(
                  (site) => site.url === thisSite
                );
                const matchedSite = sites.members[matchedSiteIndex];

                let prevSiteIndex = matchedSiteIndex - 1;
                if (prevSiteIndex === -1) prevSiteIndex = sites.members.length - 1;

                let nextSiteIndex = matchedSiteIndex + 1;
                if (nextSiteIndex > sites.members.length) nextSiteIndex = 0;

                const randomSiteIndex = this.getRandomInt(0, sites.members.length - 1);

                const cp = `
                  <h3>FTech WebRing</h3>
                  <p>
                    This <a href="${matchedSite.url}">${matchedSite.name}</a> site is owned by ${matchedSite.owner}
                  </p>
                  
                  <p>
                    <a href="${sites.members[prevSiteIndex].url}">[Prev]</a>
                    <a href="${sites.members[nextSiteIndex].url}">[Next]</a>
                    <a href="${sites.members[randomSiteIndex].url}">[Random]</a>
                  </p>
                `;

                this.shadowRoot
                  .querySelector("#copy")
                  .insertAdjacentHTML("afterbegin", cp);
            });
    }

    getRandomInt(min, max) {
        min = Math.ceil(min);
        max = Math.floor(max);
        return Math.floor(Math.random() * (max - min + 1)) + min;
    }
}

window.customElements.define("webring-css", WebRing)
