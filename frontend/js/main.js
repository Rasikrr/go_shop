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
    if (params.price) {
        var decodedPrice = decodeURIComponent(params.price[0]);
        var priceRange = decodedPrice.split('-');
        var minPrice = parseInt(priceRange[0].replace('$', ''));
        var maxPrice = parseInt(priceRange[1].replace('$', ''));
        console.log(minPrice, maxPrice)

        // Initialize the slider with the price range
        $("#slider-range").slider({
            range: true,
            min: 0,
            max: 500,
            values: [minPrice, maxPrice],
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
