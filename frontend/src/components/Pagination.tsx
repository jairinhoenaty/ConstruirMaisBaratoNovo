import React from "react";
import PropTypes from "prop-types";

interface Props {
  currentPage: number;
  totalPages: number;
  handleNextPage: (page: number) => void;
  handlePrevPage: (page: number) => void;
}
const Pagination: React.FC<Props> = ({
  currentPage,
  totalPages,
  handlePrevPage,
  handleNextPage,
}) => {
  return (
    <div className="flex justify-center items-center space-x-4 text-sm text-gray-700 mt-4">
      <button
        onClick={() => handlePrevPage(currentPage)}
        disabled={currentPage === 1}
        className={`px-3 py-1 border rounded-md shadow-sm transition-all ${
          currentPage === 1
            ? "bg-gray-200 text-gray-400 cursor-not-allowed"
            : "bg-white hover:bg-gray-100"
        }`}
      >
        &larr;
      </button>

      <span className="text-gray-600 font-medium">
        Page {currentPage} of {totalPages}
      </span>

      <button
        onClick={() => handleNextPage(currentPage)}
        disabled={currentPage === totalPages}
        className={`px-3 py-1 border rounded-md shadow-sm transition-all ${
          currentPage === totalPages
            ? "bg-gray-200 text-gray-400 cursor-not-allowed"
            : "bg-white hover:bg-gray-100"
        }`}
      >
        &rarr;
      </button>
    </div>
  );
};

Pagination.propTypes = {
  currentPage: PropTypes.number.isRequired,
  totalPages: PropTypes.number.isRequired,
  handlePrevPage: PropTypes.func.isRequired,
  handleNextPage: PropTypes.func.isRequired,
};

export default Pagination;
