import React from "react";
import Navbar from "../components/Navbar";
import SearchBar from "../components/Searchbar";
import Filters from "../components/filters";
import SearchResult from "../components/SearchResult";

const allResults = [
  {
    title: "Modern 2-Bedroom Apartment",
    image: "https://via.placeholder.com/150",
    stars: 4,
    rating: "4.5",
    price: "$120,000",
    distance: 2.5,
    userRating: 150
  },
  {
    title: "Luxury Studio Suite",
    image: "https://via.placeholder.com/150",
    stars: 5,
    rating: "5.0",
    price: "$80,000",
    distance: 1.2,
    userRating: 200
  },
  {
    title: "Spacious 3-Bedroom Villa",
    image: "https://via.placeholder.com/150",
    stars: 3,
    rating: "3.8",
    price: "$250,000",
    distance: 5.0,
    userRating: 75
  },
  {
    title: "Cozy 1-Bedroom Flat",
    image: "https://via.placeholder.com/150",
    stars: 4,
    rating: "4.2",
    price: "$95,000",
    distance: 3.0,
    userRating: 120
  },
  {
    title: "Elegant 4-Bedroom Mansion",
    image: "https://via.placeholder.com/150",
    stars: 5,
    rating: "4.8",
    price: "$500,000",
    distance: 8.0,
    userRating: 250
  }
];


const Search = () => {
  // State to hold the search query and filter values
  const [searchQuery, setSearchQuery] = useState("");
  const [filters, setFilters] = useState({
    priceRange: null,
    rating: null,
    distance: null,
  });
  const [results, setResults] = useState([]);
  


  // Function to handle search query change
  const handleSearchQueryChange = (query) => {
    setSearchQuery(query);
  };

  // Function to handle filter change
  const handleFiltersChange = (newFilters) => {
    setFilters((prevFilters) => ({
      ...prevFilters,
      ...newFilters,
    }));
  };

  // Function to handle search button click (or trigger the filter updates)
  const handleSearch = () => {
    setResults(allResults)
  };

  return (
    <div>
      <Navbar auth={true} />
      <div className="pt-20 mx-28">
        <SearchBar onSearchQueryChange={handleSearchQueryChange} onSearch={handleSearch} />
        <div className="flex pt-2">
          <Filters onFiltersChange={handleFiltersChange} onSearch={handleSearch} />
          <SearchResult results={results} />
        </div>
      </div>
    </div>
  );
};

export default Search;