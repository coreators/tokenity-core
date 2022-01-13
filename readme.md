# tokenity
**tokenity** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).


# Project Story

## Inspiration

Recently, there has been a lot of talk about new web forms such as web 3.0, so I wanted to create a decentralized social network. In a social network operating on the blockchain, each creator publish their own tokens, introduces themselves like existing SNS, and increases their value, the token value also increases.

## What it does

Tokenity is a decentralized social network platform. People can post photos and articles on our platform and publish them to nft. Creators can mint and trade their coins.

## Challenges we ran into

It was difficult to start development because the cosmos sdk had less data and a small community compared to other widely used libraries. There was no example of the function I wanted, so I had to get lost in many areas.

## Accomplishments that we're proud of

This hackathon was an opportunity to actually plan and develop a project that I had only vaguely thought of before. Also, it was good to develop using Cosmos, which I was interested in, and to learn and utilize Cosmos technology. I'm glad that I have configured the front-end in the desired shape to some extent, and I can see the direction that needs to be further developed on top of this.

## What we learned

Through this hackathon, I learned a lot about Cosmos and Tendermint and was able to develop it to some extent. Also, I was able to create the program I wanted by communicating between the web frontend and the blockchain node.

## What's next for Tokenity

Since much of the current development is unfinished, I will continue to develop even after the hackathon, and I will create functions that I have previously thought of. Also, using archway or pylons, I will issue cw20 and cw721 smart contracts to generate creator's tokens and create nft.

## Built with

What languages, frameworks, platforms, cloud services, databases, APIs, or other technologies did you use?

- Go, javascript, react, cosmos, starport

## "Try it out" links

Add links where people can try your project or see your code.

- https://github.com/coreators/tokenity-frontend
- https://github.com/coreators/tokenity-core

# Project Media

## Video demo link

This video will be embedded at the top of your project page. Read more about uploading videos.

[https://www.youtube.com/watch?v=usSGQY04Aeo](https://www.youtube.com/watch?v=usSGQY04Aeo)


# Tokenity-core

## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

## Environments
* Starport v0.19.1
* Go version go1.17 darwin/amd64

## Module
### blog
Posts in the key-value store look like this:
```
"post-0": {
  "Creator": "",
  "Images": "",
  "Content": "",
  "Replies": "",
}
```

## CLI

## API

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).

### Launch

To launch your blockchain live on multiple nodes, use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.starport.network/coreators/tokenity@latest! | sudo bash
```
`coreators/tokenity` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Starport Docs](https://docs.starport.network)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
- [Discord](https://discord.gg/cosmosnetwork)
