const templateblog = document.createElement('template');

templateblog.innerHTML = `
<style>
a {
    text-decoration: underline;
    color: #f5a500;
}
</style>
<div class="BlogWidget">
    <div id="blogcontent">

    </div>
</div>
`

class BlogWidget extends HTMLElement {
    connectedCallback() {
        this.attachShadow({ mode: 'open' });
        this.shadowRoot.appendChild(templateblog.content.cloneNode(true));

        const feed_url = this.getAttribute('url');

        fetch(feed_url)
            .then((response) => response.json())
            .then((data) => {

                let html = '<ul>'

                data.posts.forEach(post => {
                    let date = new Date(new Date(post.published_at)).toDateString();
                    html += `<li><a href="${post.url}">${date} - ${post.title}</a></li>`
                });

                html += '</ul>'

                this.shadowRoot
                    .querySelector("#blogcontent")
                    .insertAdjacentHTML("afterbegin", html);

            });
    }
}

window.customElements.define("blog-widget", BlogWidget)
