<h1 align="center">
  ML Voucher Calculator
  <br>
</h1>
<h4 align="center">Maximize total spent of a voucher according to a list of items.</h4>
<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#dependencies">Dependencies</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#documentation">Documentation</a>
</p>

## Key Features

* Maximize total spent of a voucher with N items.
* The items are represented by a list of tuples (item, price).
* The algorithm is greedy inspired by the [Knapsack problem](https://en.wikipedia.org/wiki/Knapsack_problem).

## Dependencies

* [Go (1.17) or later](https://go.dev/)
* [Google pubsub-emulator](https://cloud.google.com/pubsub/docs/emulator)
* [Mercado Libre API](https://developers.mercadolibre.com.ar/es_ar/items-y-busquedas)

## How To Use

### Running the project

Before running the project please ensure that all the dependencies are installed in your system. Then follow the next:


1. Run pubsub emulator. [Click here](https://cloud.google.com/pubsub/docs/emulator)
2. Create topic.
    - Topic: `voucher-metric`
    - Subscription: `voucher-metric-sub`
3. Run the project itself.

    ```
    make web
    ```

### Running the tests

In order to run the project tests you need to execute the following command:

```
make test
```

### TODO:
- [ ] Refactor transport layer
- [ ] Integration tests
- [ ] Circuit breaker
- [ ] Logging
- [ ] Monitoring
- [ ] Error handling
- [ ] Documentation

## Documentation



* If you want to add new features to this project please [see the contribution guide](.github/CONTRIBUTING.md)
* Questions?, <a href="mailto:machester4@gmail.com?Subject=Question about Project" target="_blank">write here</a>