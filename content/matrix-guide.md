## Jae's Guide to Matrix

This is a small guide to the chat protocol [Matrix](https://matrix.org).  
First, let's lay off some foundations.

Matrix is a federated chat protocol created in 2014 and developed by The Matrix.org Foundation.  
This protocol offers E2EE (End-to-end encrypted) communications.

To explain what federation is, let's take something we all know: e-mail.  
If you have a gmail.com address, you can still send an e-mail to your friend using protonmail.com, this is what federation is.  
Anyone can run a server and reach any user on any server federating and implementing the spec correctly.

First to start with Matrix, you will need a server. It is not recommended to use the Matrix.org server as it is overloaded and kinda centralized.  
The best option is still to run your own server if you can.

The most popular server implementations are:

 * [Synapse](https://github.com/matrix-org/synapse); the first server, written in Python
 * [Dendrite](https://github.com/matrix-org/dendrite); the next-gen server, written in Golang and set to replace Synapse in the long run
 * [Conduit](https://conduit.rs/); a server written in Rust

If you wish to host your own server, please refer to the different guides and documentation pages of the projects.

For regular users, some good servers are:

 * [TeDomum](https://element.tedomum.net); a French nonprofit (warning, pages are in french, server has a SSO to login and register)
 * [KDE](https://webchat.kde.org); the server of the KDE project (**has a monthly limit of registrations**)
 * [Tomesh](https://chat.tomesh.net); Toronto Mesh

When you have chosen your server, now let's go over the clients.  
The most popular clients are:

 * [Element](https://element.io); the 'flagship' client, made using Electron & JavaScript
 * [Nheko](https://github.com/Nheko-Reborn/nheko); a native client made using C++

For the purpose of this tutorial, we will use Element with the KDE server.  
First off, go to the homepage of Element. You can either download it or use it in your browser.  
We're gonna use it in our browser (the desktop version work the same way, it is just more convenient).

You will be greeted by this page, now click on 'Sign In'.

![Homepage of Element](https://bm.jae.fi/web/jae.fi/matrix/elementhomepage.png)

Now on this page, click the 'edit' button, select 'other homeserver' and enter the address kde.org and then validate with 'continue'. Also, don't forget to click on the 'create account' button on the bottom of the dialog.

![Edition of the homeserver](https://bm.jae.fi/web/jae.fi/matrix/matrixedit.png)

Now all you have to do is to enter an username, a password, an email address (some server might not ask for it) and click on 'Sign in'.  
You will now see the main Element interface from which you can create and join rooms.  
A first chat will be opened with a user named `riot-bot` which will display the basics of Matrix to you.

Congratulations, you are now a Matrix user!
