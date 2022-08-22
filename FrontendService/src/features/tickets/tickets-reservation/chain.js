export const formChainFor = (chainInfos, elements) => {
  const atomicActions = {
    hide: (actionInfo, type, data) => {
      elements[actionInfo.element]?.hide();
    },
    show: (actionInfo, type, data) => {
      elements[actionInfo.element]?.show();
    },
    setStyle: (actionInfo, type, data) => {
      Object.keys(actionInfo.value).forEach(key => {
        elements[actionInfo.element]?.setStyle(key, actionInfo.value[key]);
      })
    },
    setText: (actionInfo, type, data) => {
      elements[actionInfo.element]?.setText(() => interpolate(actionInfo.value, data));
    },
    goTo: (actionInfo, type, data) => {
      window.location.href = actionInfo.path;
    },
  };
  return chainInfos.map(chainInfo => {
    const forFilter = chainInfo.sensorId ? forSensor : forAny;
    return forFilter(chainInfo.sensorId,
      ...(chainInfo.onValue ? chainInfo.onValue.map(actionInfo => formAction(actionInfo, atomicActions)) : []),
      ...(chainInfo.when ? Object.keys(chainInfo.when).map(whenType => 
        formWhen(whenType, chainInfo.when, atomicActions)
      ) : []),
    )
  });
}

const formWhen = (whenType, whenInfo, atomicActions) => {
  return (type, data) => {
    Object.keys(whenInfo[whenType]).forEach(expression => {
      when(whenType, expression,
        ...whenInfo[whenType][expression].map(actionInfo => formAction(actionInfo, atomicActions))
      )(type, data);
    })
  }
}

const formAction = (actionInfo, atomicActions) => {
  return (type, data) => {
    Object.keys(actionInfo).forEach(key => {
      atomicActions[key](actionInfo[key], type, data);
    });
  }
}

// Filter by type of event and sensorId
const forSensor = (sensorId, ...successors) => (type, data) => {
  const typeConditions = {
    data: () => data?.sensor.id === sensorId,
    alarm: () => data?.sensor.id === sensorId,
    sensorStatus: () => data?.sensorId === sensorId,
    click: () => false
  }
  if(typeConditions[type]()) {
    successors?.forEach(succ => succ(type, data));
  }
}

// Call all successors
const forAny = (_, ...successors) => (type, data) => {
  successors?.forEach(succ => succ(type, data));
}

// Replaces '${nested.property.path}' to property value from object { nested: { property: { path: 1 } } } to '1'
const interpolate = (str, obj) => str.replace(/\${(.*?)}/g, (_,g) => g.split('.').reduce((o,i)=> o[i], obj));

// Filter by type of event and expression
const when = (conditionType, conditionExpression, ...successors) => (type, data) => {
  if(conditionType === type) {
    // Interpolate values and evaluate the expression
    if(eval(interpolate(conditionExpression, data))) {
      successors?.forEach(succ => succ(type, data));
    }
  }
}