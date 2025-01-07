import { useState } from "react";
import logoImg from "../../../assets/Findpeace-logo.svg";
import { FaChevronDown } from "react-icons/fa";
import Flag from "react-world-flags"; // For displaying flags

const Navbar = () => {
  const [country, setCountry] = useState("NG"); // Default country (Nigeria)
  const [language, setLanguage] = useState("EN"); // Default language
  const [currency, setCurrency] = useState("NGN"); // Default currency
  const [openDropdown, setOpenDropdown] = useState(""); // Tracks which dropdown is open

  // Dummy data for dropdowns
  const countries = [
    { name: "Nigeria", code: "NG" },
    { name: "Benin Republic", code: "BJ" }
  ];
  const languages = ["EN", "FR"];
  const currencies = ["NGN", "CFA"];

  // Handle dropdown toggle
  const toggleDropdown = dropdown => {
    setOpenDropdown(openDropdown === dropdown ? "" : dropdown); // Close if open, open if closed
  };

  // Close dropdown when clicking outside
  const handleBlur = () => {
    setTimeout(() => setOpenDropdown(""), 150); // Timeout ensures dropdown stays open during a click
  };

  return (
    <header className="px-4 md:px-[6.6rem] min-w-full h-20 fixed top-0 left-0 bg-white flex justify-between items-center z-[1000] shadow-[0px_2px_5px_rgba(0,0,0,0.1)]">
      {/* Logo */}
      <div>
        <img src={logoImg} alt="Find peace logo" />
      </div>

      {/* Flag, Language, and Currency Dropdowns */}
      <div className="hidden md:flex items-center gap-4 w-64">
        {/* Country Dropdown */}
        <div className="relative" onBlur={handleBlur} tabIndex={0}>
          <div
            className="flex items-center gap-2 cursor-pointer"
            onClick={() => toggleDropdown("country")}
          >
            <Flag code={country} className="w-6 h-4" />
            <FaChevronDown size={12} />
          </div>
          {/* Dropdown Menu */}
          {openDropdown === "country" &&
            <div className="absolute top-full left-0 bg-white shadow-[0px_4px_6px_rgba(0,0,0,0.1)] rounded-md mt-2 z-[100] animate-[fadeIn_0.2s_ease-in-out]  flex flex-col items-start gap-2.5 text-left pl-4 pt-4  w-44 h-34">
              <p>Country</p>
              {countries.map(c =>
                <div
                  key={c.code}
                  className="cursor-pointer gap-2"
                  onClick={() => {
                    setCountry(c.code);
                    setOpenDropdown("");
                  }}
                >
                  <Flag
                    code={c.code}
                    style={{ width: "24px", height: "16px" }}
                  />
                  <span
                    style={{
                      fontSize: "14px",
                      fontWeight: "500",
                      lineHeight: "19.6px"
                    }}
                  >
                    {c.name}
                  </span>
                </div>
              )}
            </div>}
        </div>

        {/* Language Dropdown */}
        <div style={{ position: "relative" }} onBlur={handleBlur} tabIndex={0}>
          <div
            style={{
              display: "flex",
              alignItems: "center",
              gap: "8px",
              cursor: "pointer",
              borderLeft: "1px solid #b2b3b3",
              borderRight: "1px solid # b2b3b3",
              paddingLeft: "8px",
              paddingRight: "10px"
            }}
            onClick={() => toggleDropdown("language")}
          >
            <span
              style={{
                fontSize: "14px",
                fontWeight: "500",
                lineHeight: "19.6px"
              }}
            >
              {language}
            </span>
            <FaChevronDown size={12} />
          </div>
          {/* Dropdown Menu */}
          {openDropdown === "language" &&
            <div
              style={{
                position: "absolute",
                top: "100%",
                left: 0,
                backgroundColor: "#ffffff",
                boxShadow: "0px 4px 6px rgba(0, 0, 0, 0.1)",
                borderRadius: "4px",
                marginTop: "8px",
                animation: "fadeIn 0.2s ease-in-out",
                zIndex: 100
              }}
            >
              {languages.map(lang =>
                <div
                  key={lang}
                  style={{
                    padding: "8px 16px",
                    cursor: "pointer"
                  }}
                  onClick={() => {
                    setLanguage(lang);
                    setOpenDropdown("");
                  }}
                >
                  {lang}
                </div>
              )}
            </div>}
        </div>

        {/* Currency Dropdown */}
        <div style={{ position: "relative" }} onBlur={handleBlur} tabIndex={0}>
          <div
            style={{
              display: "flex",
              alignItems: "center",
              gap: "8px",
              cursor: "pointer"
            }}
            onClick={() => toggleDropdown("currency")}
          >
            <span
              style={{
                fontSize: "14px",
                fontWeight: "500",
                lineHeight: "19.6px"
              }}
            >
              {currency}
            </span>
            <FaChevronDown size={12} />
          </div>
          {/* Dropdown Menu */}
          {openDropdown === "currency" &&
            <div
              style={{
                position: "absolute",
                top: "100%",
                left: 0,
                backgroundColor: "#ffffff",
                boxShadow: "0px 4px 6px rgba(0, 0, 0, 0.1)",
                borderRadius: "4px",
                marginTop: "8px",
                animation: "fadeIn 0.2s ease-in-out",
                zIndex: 100
              }}
            >
              {currencies.map(cur =>
                <div
                  key={cur}
                  style={{
                    padding: "8px 16px",
                    cursor: "pointer"
                  }}
                  onClick={() => {
                    setCurrency(cur);
                    setOpenDropdown("");
                  }}
                >
                  {cur}
                </div>
              )}
            </div>}
        </div>
      </div>
    </header>
  );
};

export default Navbar;
