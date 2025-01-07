import { useDispatch, useSelector } from "react-redux";
import { Link, useNavigate } from "react-router-dom";
import { useState } from "react";
import Navbar from "../auth/components/Navbar";
import hotels from "../../assets/imgs/hotels.png"
import Loading from "../../components/ButtonLoader";
import { LuEye, LuEyeOff } from "react-icons/lu";
import { CgDanger } from "react-icons/cg";
import googleIcon from "../../assets/google_icon.svg";
import facebookIcon from "../../assets/facebook_icon.svg";
import { HotelCard } from "./component/HotelCard";
import SearchBar from "./component/Searchbar";

import { Swiper, SwiperSlide } from "swiper/react";
import { Navigation } from "swiper/modules";
import "swiper/css";
import "swiper/css/navigation";
import { IoIosArrowForward, IoIosArrowBack } from "react-icons/io";

import { Footer } from "./component/Footer";


const Home = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
  
    const { isLoading, error } = useSelector((state) => state.auth);
  
    const handleSubmit = async (e) => {
      e.preventDefault();
  
      if (validate()) {
        dispatch(signup(formData))
          .unwrap()
          .then(() => {
            navigate("/signin");
          })
          .catch((err) => {
            console.error("Signup error:", err);
          });
      }
    };

      const popularHotels = [
        {
          name: "Joypate Hotel & Suite",
          location: "Oshodi-Isolo",
          image: "/placeholder.svg?height=400&width=600",
          rating: 4.1,
          reviews: 50,
          price: 75000
        },
        {
          name: "Eko Hotel & Suite",
          location: "Oshodi-Isolo",
          image: "/placeholder.svg?height=400&width=600",
          rating: 4.2,
          reviews: 85,
          price: 75000
        },
        {
          name: "Eko Hotel & Suite",
          location: "Oshodi-Isolo",
          image: "/placeholder.svg?height=400&width=600",
          rating: 4.2,
          reviews: 85,
          price: 75000
        },
        {
          name: "Eko Hotel & Suite",
          location: "Oshodi-Isolo",
          image: "/placeholder.svg?height=400&width=600",
          rating: 4.2,
          reviews: 85,
          price: 75000
        },
        {
          name: "Lagos Continental Hotel",
          location: "Oshodi-Isolo",
          image: "/placeholder.svg?height=400&width=600",
          rating: 4.3,
          reviews: 92,
          price: 75000
        },
        {
          name: "Golden Tulip Hotel",
          location: "Oshodi-Isolo",
          image: "/placeholder.svg?height=400&width=600",
          rating: 4.1,
          reviews: 78,
          price: 75000
        }
      ];
  
    return (
      <div className="">
        <Navbar />

        <section className="relative w-full">
          <div className="relative h-[600px] flex justify-center items-center">
            <img
              className="w-full h-[484px] object-fill "
              src={hotels}
              alt="Hero Image"
            />

            <div className="absolute inset-0 flex flex-col justify-center items-center text-white">
              <h1 className="text-4xl font-bold">
                A piece of paradise just for you
              </h1>
            </div>

            <div className="absolute mt-[450px] inset-0 flex flex-col justify-center items-center">
              <div className="text-white backdrop-blur-sm bg-white/15 w-[632px] rounded-t-xl shadow-sm border-b-white border-b">
                <ul className="flex items-center justify-center">
                  <li className="p-5">Hotels</li>
                  <li className="p-5">Flights</li>
                  <li className="p-5">Rentals</li>
                  <li className="p-5">Events</li>
                </ul>
              </div>

              <div className=" bg-white rounded-lg shadow-md w-full max-w-[1228px] flex flex-col md:flex-row gap-4">
                <div>
                  <p className="p-3">Where do you want to stay?</p>
                  <SearchBar />
                </div>
              </div>
            </div>
          </div>
        </section>

        <div className="max-w-[1228px] mx-auto">
          <section className="mt-7 rounded-t-lg">
            <div className="py-10 flex justify-center space-x-8">
              <div className="text-center">
                <img
                  src="path-to-icon1"
                  alt="Icon"
                  className="h-12 w-12 mx-auto"
                />
                <h3 className="text-lg font-bold mt-4">
                  Seamless Booking Experience
                </h3>
                <p className="text-gray-600 mt-2">
                  Our easy-to-use platform guarantees a smooth and secure
                  booking process.
                </p>
              </div>
              <div className="text-center">
                <img
                  src="path-to-icon2"
                  alt="Icon"
                  className="h-12 w-12 mx-auto"
                />
                <h3 className="text-lg font-bold mt-4">
                  Secure & Simple Payments
                </h3>
                <p className="text-gray-600 mt-2">
                  Easily manage bookings with secure payments and instant
                  updates.
                </p>
              </div>
              <div className="text-center">
                <img
                  src="path-to-icon3"
                  alt="Icon"
                  className="h-12 w-12 mx-auto"
                />
                <h3 className="text-lg font-bold mt-4">Customer Support</h3>
                <p className="text-gray-600 mt-2">
                  Assistance with bookings, payments, and questions—always here
                  to help.
                </p>
              </div>
            </div>
          </section>
        </div>

        <section className="py-16">
          <div className="container mx-auto px-4">
            <h2 className="text-2xl font-semibold mb-8">
              Popular Hotels in Nigeria
            </h2>
            {/* Custom Navigation Buttons */}
            <div className="relative">
              <button
                className=" absolute -right-5 top-1/2 transform -translate-y-1/2 bg-white text-[#5627FF] flex justify-center items-center rounded-full p-3 shadow-xl z-10"
                id="customPrev"
              >
                <IoIosArrowForward />
              </button>
              <button
                className=" absolute -left-5 top-1/2 transform -translate-y-1/2 bg-white text-[#5627FF] flex justify-center items-center rounded-full p-3 shadow-xl z-10"
                id="customNext"
              >
                <IoIosArrowBack />
              </button>

              {/* Swiper */}
              <Swiper
                modules={[Navigation]}
                navigation={{
                  prevEl: "#customPrev",
                  nextEl: "#customNext"
                }}
                spaceBetween={20}
                slidesPerView={1}
                breakpoints={{
                  640: { slidesPerView: 1 },
                  768: { slidesPerView: 2 },
                  1024: { slidesPerView: 4 }
                }}
              >
                {popularHotels.map((hotel, index) => (
                  <SwiperSlide key={index}>
                    <HotelCard {...hotel} />
                  </SwiperSlide>
                ))}
              </Swiper>
            </div>
          </div>
        </section>

        <Footer />

      </div>
    );
  };

export default Home;
