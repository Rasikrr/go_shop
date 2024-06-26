/*  ---------------------------------------------------
    Template Name: Male Fashion
    Description: Male Fashion - ecommerce teplate
    Author: Colorib
    Author URI: https://www.colorib.com/
    Version: 1.0
    Created: Colorib
---------------------------------------------------------  */

'use strict';

(function ($) {

    /*------------------
        Preloader
    --------------------*/
    $(window).on('load', function () {
        $(".loader").fadeOut();
        $("#preloder").delay(200).fadeOut("slow");

        /*------------------
            Gallery filter
        --------------------*/
        $('.filter__controls li').on('click', function () {
            $('.filter__controls li').removeClass('active');
            $(this).addClass('active');
        });
        if ($('.product__filter').length > 0) {
            var containerEl = document.querySelector('.product__filter');
            var mixer = mixitup(containerEl);
        }
    });

    /*------------------
        Background Set
    --------------------*/
    $('.set-bg').each(function () {
        var bg = $(this).data('setbg');
        $(this).css('background-image', 'url(' + bg + ')');
    });

    //Search Switch
    $('.search-switch').on('click', function () {
        $('.search-model').fadeIn(400);
    });

    $('.search-close-switch').on('click', function () {
        $('.search-model').fadeOut(400, function () {
            $('#search-input').val('');
        });
    });

    /*------------------
		Navigation
	--------------------*/
    $(".mobile-menu").slicknav({
        prependTo: '#mobile-menu-wrap',
        allowParentLinks: true
    });

    /*------------------
        Accordin Active
    --------------------*/
    $('.collapse').on('shown.bs.collapse', function () {
        $(this).prev().addClass('active');
    });


    //Canvas Menu
    $(".canvas__open").on('click', function () {
        $(".offcanvas-menu-wrapper").addClass("active");
        $(".offcanvas-menu-overlay").addClass("active");
    });

    $(".offcanvas-menu-overlay").on('click', function () {
        $(".offcanvas-menu-wrapper").removeClass("active");
        $(".offcanvas-menu-overlay").removeClass("active");
    });

    /*-----------------------
        Hero Slider
    ------------------------*/
    $(".hero__slider").owlCarousel({
        loop: true,
        margin: 0,
        items: 1,
        dots: false,
        nav: true,
        navText: ["<span class='arrow_left'><span/>", "<span class='arrow_right'><span/>"],
        animateOut: 'fadeOut',
        animateIn: 'fadeIn',
        smartSpeed: 1200,
        autoHeight: false,
        autoplay: false
    });

    /*--------------------------
        Select
    ----------------------------*/
    $("select").niceSelect();

    /*-------------------
		Radio Btn
	--------------------- */
    $(".product__color__select label, .shop__sidebar__size label, .product__details__option__size label").on('click', function (event) {
        event.stopPropagation();
        // If the label is clicked and contains a radio button, handle it properly
        var $input = $(this).find('input[type="radio"]');
        if ($input.length > 0) {
            $input.prop('checked', !$input.prop('checked'));
        }
        $(this).toggleClass('active');
        console.log('Class toggled:', $(this).attr('class'));
    });


    /*-------------------
		Scroll
	--------------------- */
    $(".nice-scroll").niceScroll({
        cursorcolor: "#0d0d0d",
        cursorwidth: "5px",
        background: "#e5e5e5",
        cursorborder: "",
        autohidemode: true,
        horizrailenabled: false
    });

    /*------------------
        CountDown
    --------------------*/
    // For demo preview start
    var today = new Date();
    var dd = String(today.getDate()).padStart(2, '0');
    var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
    var yyyy = today.getFullYear();

    if(mm == 12) {
        mm = '01';
        yyyy = yyyy + 1;
    } else {
        mm = parseInt(mm) + 1;
        mm = String(mm).padStart(2, '0');
    }
    var timerdate = mm + '/' + dd + '/' + yyyy;
    // For demo preview end


    // Uncomment below and use your date //

    /* var timerdate = "2020/12/30" */

    $("#countdown").countdown(timerdate, function (event) {
        $(this).html(event.strftime("<div class='cd-item'><span>%D</span> <p>Days</p> </div>" + "<div class='cd-item'><span>%H</span> <p>Hours</p> </div>" + "<div class='cd-item'><span>%M</span> <p>Minutes</p> </div>" + "<div class='cd-item'><span>%S</span> <p>Seconds</p> </div>"));
    });

    /*------------------
		Magnific
	--------------------*/
    $('.video-popup').magnificPopup({
        type: 'iframe'
    });

    /*-------------------
		Quantity change
	--------------------- */
    var proQty = $('.pro-qty');
    proQty.prepend('<span class="fa fa-angle-up dec qtybtn"></span>');
    proQty.append('<span class="fa fa-angle-down inc qtybtn"></span>');
    proQty.on('click', '.qtybtn', function () {
        var $button = $(this);
        var oldValue = $button.parent().find('input').val();
        if ($button.hasClass('inc')) {
            var newVal = parseFloat(oldValue) + 1;
        } else {
            // Don't allow decrementing below zero
            if (oldValue > 0) {
                var newVal = parseFloat(oldValue) - 1;
            } else {
                newVal = 0;
            }
        }
        $button.parent().find('input').val(newVal);
    });

    var proQty = $('.pro-qty-2');
    proQty.prepend('<span class="fa fa-angle-left dec qtybtn"></span>');
    proQty.append('<span class="fa fa-angle-right inc qtybtn"></span>');
    proQty.on('click', '.qtybtn', function () {
        var $button = $(this);
        var oldValue = $button.parent().find('input').val();
        if ($button.hasClass('inc')) {
            var newVal = parseFloat(oldValue) + 1;
        } else {
            // Don't allow decrementing below zero
            if (oldValue > 0) {
                var newVal = parseFloat(oldValue) - 1;
            } else {
                newVal = 0;
            }
        }
        $button.parent().find('input').val(newVal);
    });

    /*------------------
        Achieve Counter
    --------------------*/
    $('.cn_num').each(function () {
        $(this).prop('Counter', 0).animate({
            Counter: $(this).text()
        }, {
            duration: 4000,
            easing: 'swing',
            step: function (now) {
                $(this).text(Math.ceil(now));
            }
        });
    });

})(jQuery);

$(function() {
    $("#slider-range").slider({
        range: true,
        min: 0,
        max: 500,
        values: [0, 500],
        slide: function(event, ui) {
            $("#amount").val("$" + ui.values[0] + " - $" + ui.values[1]);
        }
    });
    $("#amount").val("$" + $("#slider-range").slider("values", 0) +
        " - $" + $("#slider-range").slider("values", 1));
});



$(document).ready(function() {
    var params = getQueryParams();

    // Check radio buttons and checkboxes based on the parsed query parameters
    if (params.sex) {
        params.sex.forEach(function(value) {
            $('input[name="sex"][value="' + value + '"]').prop('checked', true);
        });
    }

    if (params.categories) {
        params.categories.forEach(function(value) {
            $('input[name="categories"][value="' + value + '"]').prop('checked', true);
        });
    }

    if (params.brands) {
        params.brands.forEach(function(value) {
            $('input[name="brands"][value="' + value + '"]').prop('checked', true);
        });
    }

    if (params.sizes) {
        params.sizes.forEach(function(value) {
            $('input[name="sizes"][value="' + value + '"]').prop('checked', true);
        });
    }
    var minVal = parseInt(document.getElementById('amount').getAttribute('min'));
    var maxVal = parseInt(document.getElementById('amount').getAttribute('max'));
    if (params.price) {
        var decodedPrice = decodeURIComponent(params.price[0]);
        var priceRange = decodedPrice.split('-');
        var minPrice = parseInt(priceRange[0].replace('$', ''));
        var maxPrice = parseInt(priceRange[1].replace('$', ''));
        console.log(minPrice, maxPrice)


        // Initialize the slider with the price range
        $("#slider-range").slider({
            range: true,
            min: minVal,
            max: maxVal,
            values: [minPrice, maxPrice],
            slide: function(event, ui) {
                $("#amount").val("$" + ui.values[0] + " - $" + ui.values[1]);
            }
        });
        $("#amount").val("$" + $("#slider-range").slider("values", 0) +
            " - $" + $("#slider-range").slider("values", 1));
    } else{
        $("#slider-range").slider({
            range: true,
            min: minVal,
            max: maxVal,
            values: [minVal, maxVal],
            slide: function(event, ui) {
                $("#amount").val("$" + ui.values[0] + " - $" + ui.values[1]);
            }
        });
        $("#amount").val("$" + $("#slider-range").slider("values", 0) +
            " - $" + $("#slider-range").slider("values", 1));
    }
});


function getQueryParams() {
    var queryString = window.location.search;
    var urlParams = new URLSearchParams(queryString);
    var params = {};

    for (let [key, value] of urlParams.entries()) {
        if (!params[key]) {
            params[key] = [];
        }
        params[key].push(value);
    }

    return params;
}

document.querySelectorAll('.nav-link').forEach(function(tab) {
    tab.addEventListener('click', function(e) {
        e.preventDefault();
        // Remove active class from all tabs
        document.querySelectorAll('.nav-link.active').forEach(function(activeTab) {
            activeTab.classList.remove('active');
        });
        // Remove active class from all tab content
        document.querySelectorAll('.tab-pane.active').forEach(function(activePane) {
            activePane.classList.remove('active');
        });
        // Add active class to clicked tab
        this.classList.add('active');
        // Add active class to corresponding tab content
        var id = this.getAttribute('href');
        document.querySelector(id).classList.add('active');
    });
});


function handleClick(clickedRadio) {
    // Get all radio buttons
    var radios = document.querySelectorAll('.size-radio');

    // Loop through all radio buttons
    for (var i = 0; i < radios.length; i++) {
        var radio = radios[i];
        var label = radio.parentElement; // Get the parent element (label) of the radio button

        // If the radio button is the one that was clicked, add the 'active' class to its label
        if (radio === clickedRadio) {
            label.classList.add('active');
        }
        // If the radio button is not the one that was clicked, remove the 'active' class from its label
        else {
            label.classList.remove('active');
        }
    }
}

$(document).ready(function(){
    $(".category-checkbox").change(function(){
        $(this).parent().next('ul').find('.subcategory-checkbox').prop('checked', this.checked);
    });
});

window.onload = function() {
    // Get URL parameters
    var params = new URLSearchParams(window.location.search);

    // Get all checkboxes
    var checkboxes = document.querySelectorAll('.subcategory-checkbox');

    // Loop through checkboxes
    checkboxes.forEach(function(checkbox) {
        // Get the value of the checkbox
        var value = checkbox.value;

        // Check if the URL parameters include the checkbox's value
        if (params.getAll('subcategories').includes(value)) {
            // If they do, check the checkbox
            checkbox.checked = true;
        }
    });
};
window.onload = function() {
    // Get URL parameters
    var params = new URLSearchParams(window.location.search);

    // Get all checkboxes for sizes
    var checkboxes = document.querySelectorAll('input[name="sizes"]');

    // Loop through checkboxes
    checkboxes.forEach(function(checkbox) {
        // Get the value of the checkbox
        var value = checkbox.value;

        // Check if the URL parameters include the checkbox's value
        if (params.getAll('sizes').includes(value)) {
            // If they do, check the checkbox
            checkbox.checked = true;

            // Add 'active' class to parent label element
            checkbox.parentElement.classList.add('active');
        }
    });
};
// Get all checkboxes for sizes
var checkboxes = document.querySelectorAll('input[name="sizes"]');

// Loop through checkboxes
checkboxes.forEach(function(checkbox) {
    // Add event listener to checkbox
    checkbox.addEventListener('change', function() {
        // If checkbox is checked
        if (this.checked) {
            // Add 'active' class to parent label element
            this.parentElement.classList.add('active');
        } else {
            // Remove 'active' class from parent label element
            this.parentElement.classList.remove('active');
        }
    });
});



function upload() {

    const fileUploadInput = document.querySelector('.file-uploader');

    // using index [0] to take the first file from the array
    const image = fileUploadInput.files[0];

    // check if the file selected is not an image file
    if (!image.type.includes('image')) {
        return alert('Only images are allowed!');
    }

    // check if size (in bytes) exceeds 10 MB
    if (image.size > 10_000_000) {
        return alert('Maximum upload size is 10MB!');
    }

    const fileReader = new FileReader();
    fileReader.readAsDataURL(image);

    fileReader.onload = (fileReaderEvent) => {
        const profilePicture = document.querySelector('.profile-picture');
        profilePicture.style.backgroundImage = `url(${fileReaderEvent.target.result})`;
    }
}


document.addEventListener("DOMContentLoaded", function() {
    let notificationItem = document.getElementById("notification");
    let goodNotificationItem = document.getElementById("notification-good");
    let addToCartForm = document.getElementById('addToCartForm');
    let overall = document.getElementById("cart-overall");
    // Attach the submit event to the form
    if (addToCartForm) {
        addToCartForm.addEventListener('submit', function(e) {
            // Prevent the form from submitting normally
            e.preventDefault();

            // Get the product ID, size, and amount from the form
            const productID = this.querySelector('button[product-id]').getAttribute('product-id');
            const sizeElement = this.querySelector('input[name="size"]:checked');

// Check if a size is selected
            if (!sizeElement) {
                console.error('No size selected');
                return;
            }

// Get the size from the id attribute instead of the value attribute
            const size = sizeElement.id;
            const amount = 1; // Set the amount
            console.log(size);

            // Create a CartItem object
            const cartItem = {
                ProductID: productID,
                Size: size,
                Amount: amount
            };

            // Send a POST request to the /products/add endpoint with the CartItem object in the request body
            fetch('http://localhost:8080/products/add', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(cartItem),
            })
                .then(response => response.json())
                .then(data => {
                    if (data.error !== 'ok') {
                        notificationItem.style.display = 'block'
                        document.querySelector(".notifications__item__message").innerHTML = "You have to be authorized to add items to cart"
                        console.error('Error:', data.error);
                    } else {
                        let priceElement = document.querySelector('#prod-price');
                        let priceText = priceElement.textContent;
                        let prices = priceText.split(' ');
                        let finalPrice = parseFloat(prices[0].substring(1));

                        let overallCartPrice = parseFloat(overall.innerHTML.substring(1)); // Remove the dollar sign before parsing
                        overallCartPrice += finalPrice;
                        overall.innerHTML = "$" + overallCartPrice.toFixed(2); // Add the dollar sign back when displaying the price

                        goodNotificationItem.style.display = 'block'
                        document.getElementById("good-mess").innerHTML = "Added to cart successfully"
                        console.log('Added to cart successfully');
                    }
                })
                .catch((error) => {
                    console.error('Error:', error);
                });
        });
    }
});


(function(){

    /*
    * Get all the buttons actions
    */
    let optionBtns = document.querySelectorAll( '.js-option' );

    for(var i = 0; i < optionBtns.length; i++ ) {

        /*
        * When click to a button
        */
        optionBtns[i].addEventListener( 'click', function ( e ){

            var notificationCard = this.parentNode.parentNode;
            var clickBtn = this;
            /*
            * Execute the delete or Archive animation
            */
            requestAnimationFrame( function(){

                archiveOrDelete( clickBtn, notificationCard );

                /*
                * Add transition
                * That smoothly remove the blank space
                * Leaves by the deleted notification card
                */
                window.setTimeout( function( ){
                    requestAnimationFrame( function() {
                        notificationCard.style.transition = 'all .4s ease';
                        notificationCard.style.height = 0;
                        notificationCard.style.margin = 0;
                        notificationCard.style.padding = 0;
                    });

                    /*
                    * Delete definitely the animation card
                    */
                    // window.setTimeout( function( ){
                    //     notificationCard.parentNode.removeChild( notificationCard );
                    // }, 1500 );
                }, 1500 );
            });
        })
    }

    /*
    * Function that adds
    * delete or archive class
    * To a notification card
    */
    var archiveOrDelete = function( clickBtn, notificationCard ){
        if( clickBtn.classList.contains( 'auth-btn' ) ){
            notificationCard.classList.add( 'archive' );
            window.location = "/auth"
        } else if( clickBtn.classList.contains( 'delete' ) ){
            // notificationCard.classList.add( 'delete' );
            notificationCard.style.display = 'none';

        }
        notificationCard.style.display = 'none';

    }

})()
let goodNotificationBtn = document.getElementById("good-btn")

goodNotificationBtn.addEventListener('click', function (){
    let goodNotificationItem = document.getElementById("notification-good");
    goodNotificationItem.classList.add( 'archive' );
})


gsap.set("svg", { visibility: "visible" });
gsap.to("#headStripe", {
    y: 0.5,
    rotation: 1,
    yoyo: true,
    repeat: -1,
    ease: "sine.inOut",
    duration: 1
});
gsap.to("#spaceman", {
    y: 0.5,
    rotation: 1,
    yoyo: true,
    repeat: -1,
    ease: "sine.inOut",
    duration: 1
});
gsap.to("#craterSmall", {
    x: -3,
    yoyo: true,
    repeat: -1,
    duration: 1,
    ease: "sine.inOut"
});
gsap.to("#craterBig", {
    x: 3,
    yoyo: true,
    repeat: -1,
    duration: 1,
    ease: "sine.inOut"
});
gsap.to("#planet", {
    rotation: -2,
    yoyo: true,
    repeat: -1,
    duration: 1,
    ease: "sine.inOut",
    transformOrigin: "50% 50%"
});

gsap.to("#starsBig g", {
    rotation: "random(-30,30)",
    transformOrigin: "50% 50%",
    yoyo: true,
    repeat: -1,
    ease: "sine.inOut"
});
gsap.fromTo(
    "#starsSmall g",
    { scale: 0, transformOrigin: "50% 50%" },
    { scale: 1, transformOrigin: "50% 50%", yoyo: true, repeat: -1, stagger: 0.1 }
);
gsap.to("#circlesSmall circle", {
    y: -4,
    yoyo: true,
    duration: 1,
    ease: "sine.inOut",
    repeat: -1
});
gsap.to("#circlesBig circle", {
    y: -2,
    yoyo: true,
    duration: 1,
    ease: "sine.inOut",
    repeat: -1
});

gsap.set("#glassShine", { x: -68 });

gsap.to("#glassShine", {
    x: 80,
    duration: 2,
    rotation: -30,
    ease: "expo.inOut",
    transformOrigin: "50% 50%",
    repeat: -1,
    repeatDelay: 8,
    delay: 2
});

const burger = document.querySelector('.burger');
const nav = document.querySelector('nav');

burger.addEventListener('click',(e) => {
    burger.dataset.state === 'closed' ? burger.dataset.state = "open" : burger.dataset.state = "closed"
    nav.dataset.state === "closed" ? nav.dataset.state = "open" : nav.dataset.state = "closed"
})