class PaymentChoiceForm extends KDView
  constructor: (options = {}, data) ->
    options.cssClass = KD.utils.curry "pricing-payment-choice clearfix", options.cssClass
    super options, data


  activate: (activator) -> @emit 'Activated', activator

  setPaymentMethods: (paymentMethods) ->
    @paymentMethodsContainer.addSubView new KDCustomHTMLView
      cssClass : "new payment-method"
      partial  : "Add a Payment method"
      click    : @lazyBound "emit", "PaymentMethodNotChosen"

    { preferredPaymentMethod, methods, appStorage } = paymentMethods
    for method in methods
      @paymentMethodsContainer.addSubView view = new PaymentMethodView null, method
      @forwardEvent view, "PaymentMethodChosen"

    return this

  viewAppended: ->
    @addSubView new KDCustomHTMLView
      tagName  : "h3"
      cssClass : "pricing-title"
      partial  : "Choose a payment method"

    @addSubView new KDCustomHTMLView
      tagName  : "h6"
      cssClass : "pricing-subtitle"
      partial  : "Or add a new one, whatever works"

    @addSubView @paymentMethodsContainer = new KDCustomHTMLView
      cssClass  : "payment-methods"
