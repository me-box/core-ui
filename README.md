# Databox core UI

The official databox user interface. It is started by the container manager and runs with the same permissions as an any databox app.
It uses the interface exposed through container-manager-core-store to get system information and invoke container manager API functions.

## Testing/developing outside databox

Its passable to test this component outside of databox, to do so run:

```
npm run serve
```

You will also need to set the dev flag in you browsers local storage. In the Javascript console:
```
localStorage.setItem("dev","true")
```

In this mode the API call will return canned data stored in then ./testData directory.

# TODO

- Add the ability to add new manifest stores

## Development of databox was supported by the following funding

```
EP/N028260/1, Databox: Privacy-Aware Infrastructure for Managing Personal Data

EP/N028260/2, Databox: Privacy-Aware Infrastructure for Managing Personal Data

EP/N014243/1, Future Everyday Interaction with the Autonomous Internet of Things

EP/M001636/1, Privacy-by-Design: Building Accountability into the Internet of Things (IoTDatabox)

EP/M02315X/1, From Human Data to Personal Experience

```