<div id="top"></div>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<!-- [![Contributors][contributors-shield]][contributors-url] -->
<!-- [![Forks][forks-shield]](https://img.shields.io/github/forks/leopku/meilisearch-prompt?style=for-the-badge) -->
<!-- [![Stargazers][stars-shield]](https://img.shields.io/github/stars/leopku/meilisearch-prompt?style=for-the-badge) -->
<!-- [![Issues][issues-shield]][issues-url] -->
<!-- [![AGPL License][license-shield]](https://www.gnu.org/licenses/agpl-3.0.txt) -->
<!-- [![LinkedIn][linkedin-shield]][linkedin-url] -->
<!-- ![Twitter URL](https://img.shields.io/twitter/url?style=for-the-badge&url=https%3A%2F%2Ftwitter.com%2Fleopku) -->

<!-- ![forks](https://img.shields.io/github/forks/leopku/meilisearch-prompt?style=for-the-badge) -->
<!-- ![starts](https://img.shields.io/github/stars/leopku/meilisearch-prompt?style=for-the-badge) -->
<!-- [AGPL License](https://www.gnu.org/licenses/agpl-3.0.txt) -->



<!-- PROJECT LOGO -->
<br />
<!-- <div align="center">
  <a href="https://github.com/leopku/meilisearch-prompt">
    <img src="https://s3.bmp.ovh/imgs/2022/04/17/94abaf334c384dbc.gif" alt="Logo" width="800" height="400">
  </a> -->

<h3 align="center">meilisearch prompt</h3>

  <p align="center">
    A command-line kit to manage a meilisearch server easily.
    <br />
    <a href="https://github.com/leopku/meilisearch-prompt"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <!-- <a href="https://github.com/leopku/meilisearch-prompt">View Demo</a>
    · -->
    <a href="https://github.com/leopku/meilisearch-prompt/issues">Report Bug</a>
    ·
    <a href="https://github.com/leopku/meilisearch-prompt/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

![product-screenshot](https://s3.bmp.ovh/imgs/2022/04/17/94abaf334c384dbc.gif)

A command-line kit to manage a meilisearch server easily.

<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [Golang](https://go.dev/)
* [Cobra](https://cobra.dev/)
* [go-prompt](https://pkg.go.dev/github.com/c-bata/go-prompt)
* [meilisearch](https://www.meilisearch.com/)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

A command-line kit to manage a meilisearch server with promption commands, indexes and field of one index.

### Prerequisites

* meilisearch v0.26.0

### Installation

1. A meilisearch instance of v0.26.0
2. Download binary from ![github release](https://github.com/leopku/meilisearch-prompt/releases)
3. Running `meilisearch-prompt` with interactive command-line mode.
   ```sh
   ./meilisearch-prompt interactive <your meilisearch host>
   ```
4. Manage your meilisearch instance
   ```
   ls
   cd <your index>
   settings // get all settings of current index
   settings ranking-rules //get ranking-rules settings of current index
   settings ranking-rules words typo sort // set ranking-rules settings of current index
   ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

* ls - ls all indexes of current instance
* cd <uid> - select an index as current index
* create <uid> [primaryKey] - create a new index
* info - get info about current index
* settings - get all settings of current index
  * settings <item> - get one item of current index settings.
  * settings <item> [field1] [field2] [field2] - set one item of current index settings.
> item may be one of below:
>
> displayed-attributes, searchable-attributes, filterable-attributes, sortable-attributes, ranking-rules, stop-words, distinct-attribute

<!-- _For more examples, please refer to the [Documentation](https://example.com)_ -->

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [x] ls
- [x] cd
- [x] create
- [x] settings
    - [x] displayed-attributes
      - [x] get
      - [x] set
    - [x] searchable-attributes
      - [x] get
      - [x] set
    - [x] filterable-attributes
      - [x] get
      - [x] set
    - [x] sortable-attributes
      - [x] get
      - [x] set
    - [x] ranking-rules
      - [x] get
      - [x] set
    - [x] stop-words
      - [x] get
      - [x] set
    - [x] distinct-attribute
      - [x] get
      - [x] set

See the [open issues](https://github.com/leopku/meilisearch-prompt/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- LICENSE -->
## License

Distributed under the AGPL v3 License. See `LICENSE` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

[@leopku](https://github.com/leopku)

[https://github.com/leopku/meilisearch-prompt](https://github.com/leopku/meilisearch-prompt)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

* [Meilisearch Documentation](https://docs.meilisearch.com/)
* [go package of meilisearch](https://pkg.go.dev/github.com/meilisearch/meilisearch-go)
* [go package of prompt](https://pkg.go.dev/github.com/c-bata/go-prompt)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/leopku/meilisearch-prompt.svg?style=for-the-badge
[contributors-url]: https://github.com/leopku/meilisearch-prompt/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/leopku/meilisearch-prompt.svg?style=for-the-badge
[forks-url]: https://github.com/leopku/meilisearch-prompt/network/members
[stars-shield]: https://img.shields.io/github/stars/leopku/meilisearch-prompt.svg?style=for-the-badge
[stars-url]: https://github.com/leopku/meilisearch-prompt/stargazers
[issues-shield]: https://img.shields.io/github/issues/leopku/meilisearch-prompt.svg?style=for-the-badge
[issues-url]: https://github.com/leopku/meilisearch-prompt/issues
[license-shield]: https://img.shields.io/github/license/leopku/meilisearch-prompt.svg?style=for-the-badge
[license-url]: https://github.com/leopku/meilisearch-prompt/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png