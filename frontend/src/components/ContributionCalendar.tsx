import { Box, Typography, Grid, Tooltip } from '@mui/material';

type MonthInfo = { firstDay: number; daysInMonth: number };
function getMonthInfo(year: number, month: number): MonthInfo {
  const firstDay = new Date(year, month, 1).getDay();   //!0 represent sunday.
  console.log("First day of the month: ", firstDay, "Input month: ", month, "Input year: ", year);
  //!month + 1 with day 0 gives last day of previous month, i.e., last day of current month
  const daysInMonth = new Date(year, month + 1, 0).getDate();
  console.log("Days in the month: ", daysInMonth);0
  return { firstDay, daysInMonth };
}


//!Set this color to the theme for the ease.
// Color scale (customize as needed)
const getColor = (count: number) => {
  if (count === 0) return 'pink';
  if (count === 1) return '#81c784';
  if (count >= 2) return '#4caf50';
};

//0 is January, 6 is July
export default function MonthlyContributionCalendar({ year = 2025, month = 4 }: { year?: number; month?: number }) {

  //!0 is Sunday and 6 is Saturday
  const { firstDay: firstDay1, daysInMonth: daysInMonth } = getMonthInfo(year, month);
  console.log(firstDay1, daysInMonth);


  //!First day will tell us the index form which we have to start, showing the boxes.

  //!This has to be fetched from the databse, isn't it?
  // Example data: array of { date: 'YYYY-MM-DD', count: number }
  const activityData = Array.from({ length: daysInMonth }, (_, i) => ({
    date: new Date(year, month, i + 1).toISOString().slice(0, 10),
    count: Math.floor(Math.random() * 5), // Random activity count
  }));

  
  const weeks: ({ date: string; count: number; dayOfWeek: number } | null)[][] = [];

  let currDay = 1;
  let week: ({ date: string; count: number; dayOfWeek: number } | null)[] = [];
  
  // Get the day index (0 = Sunday, ..., 6 = Saturday) for the 1st of the month
  const firstDay = new Date(year, month, 1).getDay();
  
  // Fill initial nulls for days before the first day of the month
  for (let i = 0; i < firstDay; i++) {
    week.push(null);
  }
  
  // Loop through all days in the month
  while (currDay <= daysInMonth) {
    const currDate = new Date(year, month, currDay);
    const dayOfWeek = currDate.getDay();
  
    week.push({
      date: currDate.toDateString(),
      count: 0,
      dayOfWeek: dayOfWeek
    });
  
    // If week is complete, push and reset
    if (week.length === 7) {
      weeks.push(week);
      week = [];
    }
  
    currDay++;
  }
  
  // Push the last partial week (if any)
  if (week.length > 0) {
    while (week.length < 7) {
      week.push(null);
    }
    weeks.push(week);
  }
  
  console.log(weeks);
  


  /*
  //!Need to transform it --> Right now, each row is a week. We need to make it that each column is a week.
  const transformedWeeks = weeks[0].map((_, i) => weeks.map(week => week[i]));
  
  //!Every row of the transformedWeeks belongs to the same day like -- Sunday, Monday, Tuesday, etc.
  //!We need to sort it so that the first row is always the Sunday.
  transformedWeeks.sort((a, b) => {
    let aDay = a.find(day => day !== null)?.dayOfWeek;
    let bDay = b.find(day => day !== null)?.dayOfWeek;
    return aDay! - bDay!;
  });

  */

  const transformedWeeks = weeks;

  //!Now, we need to make it that each column is a week.


  //!Transforming it to the format that we want./
  console.log(transformedWeeks);

  return (
    <Box sx={{ width: '100%', maxWidth: '100%' }}>
      <Box sx={{ display: 'flex', gap: 0.5 }}>
        {transformedWeeks.map((week, wi) => (
          <Box key={wi} sx={{ display: 'flex', flexDirection: 'column', gap: 0.5 }}>
            {week.map((day, di) => (
              <Box key={day?.dayOfWeek}>
                {day?.date ? (
                  <Tooltip title={`${day.date}: ${day.count} activities`} arrow>
                    <Box
                      sx={{
                        width: { md: '10px'},
                        height: { md: '10px'},
                        minWidth: { md: '10px'},
                        minHeight: { md: '10px'},
                        bgcolor: getColor(day.count),
                        borderRadius: 1,
                        border: '1px solid #222',
                        cursor: 'pointer',
                        transition: '0.2s',
                      }}
                    />
                  </Tooltip>
                ) : (
                  <Box 
                    sx={{ 
                      width: { md: '10px'},
                      height: { md: '10px'},
                      minWidth: { md: '10px'},
                      minHeight: { md: '10px'},
                      bgcolor: 'transparent' 
                    }} 
                  />
                )}
              </Box>
            ))}
          </Box>
        ))}
      </Box>
    </Box>
  );
}
  