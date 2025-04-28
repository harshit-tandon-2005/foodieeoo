# Shopping Cart

Build a mini food ordering web app featuring product listing and a functional shopping cart.\
Prioritize correctness in functionality while getting it to look as close to the design as possible.

For this task you will need to integrate to our demo e-commerce API for listing products and placing orders.

**API Reference**

You can find our [API Documentation](https://orderfoodonline.deno.dev/public/openapi.html) here.

API documentation is based on [OpenAPI3.1](https://swagger.io/specification/v3/) specification.
You can also find spec file [here](https://orderfoodonline.deno.dev/public/openapi.yaml).
 
**Functional Requirements**

- Display products with images
- Add items to the cart and remove items
- Show order total correctly
- Increase or decrease item count in the cart
- Show order confirmation after placing the order
- Interactive hover and focus states for elements

**Bonus Goals**

- Allow users to enter a discount code (above the "Confirm Order" button)
- Discount code `HAPPYHOURS` applies 18% discount to the order total
- Discount code `BUYGETONE` gives the lowest priced item for free
- Responsive design based on device's screen size

**Are You a Full Stack Developer??**

Impress us by implementing your own version of the API based on the OpenAPI specification.\
Choose any language or framework of your choice. For example our top pick for backend is [Go](https://go.dev)

> The API immplementation example available to you at orderfoodonline.deno.dev/api is simplified and doesn't handle some edge cases intentionally.
> Use your best judgement to build a Robust API server.

## Design

You can find a [Figma](https://figma.com) design file `design.fig` that you can use.
You might have to use your best judgement for some mobile layout designs and spacing.

### Style Guide

The designs were created to the following widths:

- Mobile: 375px
- Desktop: 1440px

> 💡 These are just the design sizes. Ensure content is responsive and meets WCAG requirements by testing the full range of screen sizes from 320px to large screens.

**Typography**

- Font size (product names): 16px

### Font

- Family: [Red Hat Text](https://fonts.google.com/specimen/Red+Hat+Text)
- Weights: 400, 600, 700

## Getting Started

Feel free to use any tool or workflow ou are comformtable with.\
Here is an example workflow (you can use it as a reference or use your own workflow)

1. Create a new public repository on [GitHub](https://github.com) (alternatively you can use GitLab, BitBucket or Git server of your choice).
   If you are creating your repository on GitHub, you can chose to use this repository as a starting template. (Click on Use template button at the top)
2. Look through the deisngs to plan your project. This will help you design UI libraries or tools.
3. Create a [Vite](https://vite.dev) app to bootstrap a modern front-end project (alternatively use the framework of your choice).
4. Structure your HTML and preview before theming and adding interactive functionality.
5. Test and Iterate to build more features
6. Deploy your app anywhere securely. You may use AWS, Vercel, Deno Deploy, Surge, CloudFlare Pages or some other web app deployment services.
7. Additionally configure your repository to automatically publish your app on new commit push (CI).

> 💡 Replace or Modify this README to explain your solution and how to run and test it.

_By following these guidelines, you should be able to build a functional and visually appealing mini e-commerce shopping portal that meets the minimum requirements and bonus goals. Good luck! 🚀_

**Resources**

- API documentation: https://orderfoodonline.deno.dev/public/openapi.html
- API specification: https://orderfoodonline.deno.dev/public/openapi.yaml
- Figma design file: [design.fig](./design.fig)
- Red Hat Text font: https://fonts.google.com/specimen/Red+Hat+Text





## foodieeoo

foodieeoo is a Food Ordering Web server application written in GoLang.


## Prerequisites

*   [Go](https://golang.org/dl/) installed (version 1.16 or later recommended).

## Setup

1.  **Clone the repository (if you haven't already):**
    ```bash
    git clone <your-repository-url>
    cd <repository-directory>
    ```

2.  **Install Dependencies:**
    To let Go manage dependencies based on imports, run:
    ```bash
    go mod tidy
    ```
    This command ensures `go.mod` file matches the packages used in the code and downloads them.

3.  **(Optional) Vendor Dependencies:**
    To include dependencies directly in your project repository for offline builds, run:
    ```bash
    go mod vendor
    ```
    This will create a `vendor` directory containing all necessary packages.

4.  **Configure:**
    Create a `config.yml` from `sample_config.yml` file and set the desired configuration values.

## Running the Script

To run the main script, execute:

```bash
go run main.go
```

If you vendored dependencies (Step 3), you might need to build or run using the `-mod=vendor` flag, although `go run` often detects the vendor directory automatically:

```bash
go run -mod=vendor main.go
# or build first
go build -mod=vendor
./<executable-name> # e.g., ./transaction-tracker
```


## Running the DB Migrations

1. Install goose to handle migrations for the service (https://github.com/pressly/goose)

2. To create a new Migration file, run the following command:

    ```goose -dir db/migrations create create_new_table sql```

3. To check the status of migrations, run the following command:

   ```goose -dir db/migrations mysql "<user>:<password>@tcp(<host>:<port>)/<database_name>?parseTime=true" status```

4. To upgrade the DB, run the following command:

   ```goose -dir db/migrations mysql "<user>:<password>@tcp(<host>:<port>)/<database_name>?parseTime=true" up```

5. To downgrade the DB, run the following command:

   ```goose -dir db/migrations mysql "<user>:<password>@tcp(<host>:<port>)/<database_name>?parseTime=true" down```

